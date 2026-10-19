package limiter

import (
	"container/list"
	"sync"
	"time"
)

// rateLimitRequest stores one accepted request in a key's sliding window.
type rateLimitRequest struct {
	next      *rateLimitRequest
	timestamp int64
}

// rateLimitQueue keeps accepted requests ordered from oldest to newest.
type rateLimitQueue struct {
	head   *rateLimitRequest
	tail   *rateLimitRequest
	length int
}

// append adds an accepted request without preallocating for the configured limit.
func (q *rateLimitQueue) append(timestamp int64) {
	request := &rateLimitRequest{timestamp: timestamp}
	if q.tail == nil {
		q.head = request
		q.tail = request
	} else {
		q.tail.next = request
		q.tail = request
	}
	q.length++
}

// removeExpired releases every expired request at the front of the queue.
func (q *rateLimitQueue) removeExpired(now int64, duration int64) {
	if q.head == nil || now-q.head.timestamp < duration {
		return
	}

	// Requests are time ordered, so an expired tail means the whole queue expired.
	if now-q.tail.timestamp >= duration {
		q.clear()
		return
	}

	for now-q.head.timestamp >= duration {
		expired := q.head
		q.head = expired.next
		expired.next = nil
		q.length--
	}
}

// clear releases the whole request chain when its key expires.
func (q *rateLimitQueue) clear() {
	q.head = nil
	q.tail = nil
	q.length = 0
}

// rateLimitEntry is both a rate-limit bucket and a node in the key-level LRU.
type rateLimitEntry struct {
	lastActive time.Time
	element    *list.Element
	requests   rateLimitQueue
	key        string
}

// InMemoryRateLimiter 滑动窗口限流器，带空闲 key 的 LRU 淘汰。
// 在途请求通过 Reserve 单独计数，key 被淘汰时不影响未完成的请求。
type InMemoryRateLimiter struct {
	store              map[string]*rateLimitEntry
	lru                *list.List
	reservations       map[string]int
	mutex              sync.Mutex
	expirationDuration time.Duration
}

// New 创建限流器。expirationDuration > 0 时启动后台 goroutine 定期清理空闲 key（0 表示不淘汰）。
func New(expirationDuration time.Duration) *InMemoryRateLimiter {
	l := &InMemoryRateLimiter{
		store:              make(map[string]*rateLimitEntry),
		lru:                list.New(),
		reservations:       make(map[string]int),
		expirationDuration: expirationDuration,
	}
	if expirationDuration > 0 {
		go l.clearExpiredItems(time.NewTicker(expirationDuration).C)
	}
	return l
}

// clearExpiredItems periodically removes expired entries from the LRU tail.
func (l *InMemoryRateLimiter) clearExpiredItems(ticks <-chan time.Time) {
	for now := range ticks {
		l.deleteExpiredEntries(now)
	}
}

// deleteExpiredEntries walks only the oldest LRU entries and stops at the first active key.
func (l *InMemoryRateLimiter) deleteExpiredEntries(now time.Time) {
	if l.expirationDuration <= 0 {
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	for {
		oldest := l.lru.Back()
		if oldest == nil {
			return
		}

		entry := oldest.Value.(*rateLimitEntry)
		if now.Sub(entry.lastActive) < l.expirationDuration {
			return
		}

		delete(l.store, entry.key)
		l.lru.Remove(oldest)
		entry.element = nil
		entry.requests.clear()
	}
}

// touchEntry returns the key's bucket, creating it on demand, and marks it as
// the most recently active key. The caller must hold the mutex.
func (l *InMemoryRateLimiter) touchEntry(key string, now time.Time) *rateLimitEntry {
	entry, ok := l.store[key]
	if !ok {
		entry = &rateLimitEntry{
			key:        key,
			lastActive: now,
		}
		entry.element = l.lru.PushFront(entry)
		l.store[key] = entry
		return entry
	}
	entry.lastActive = now
	l.lru.MoveToFront(entry.element)
	return entry
}

// Request 滑动窗口准入：key 在 duration（秒）内已接受的请求数小于 maxRequestNum 时放行并计数。
func (l *InMemoryRateLimiter) Request(key string, maxRequestNum int, duration int64) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	now := time.Now()
	entry := l.touchEntry(key, now)
	entry.requests.removeExpired(now.Unix(), duration)

	allowed := entry.requests.length < maxRequestNum
	if allowed {
		entry.requests.append(now.Unix())
	}

	return allowed
}

// Reservation 在途请求计数：成功完成时记入滑动窗口，失败时仅释放名额
type Reservation struct {
	limiter *InMemoryRateLimiter
	key     string
	once    sync.Once
}

// Reserve 在途请求数 + 窗口内请求数 < maxRequests 时放行，否则返回 nil。duration 单位为秒。
func (l *InMemoryRateLimiter) Reserve(key string, maxRequests int, duration int64) *Reservation {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	now := time.Now()
	entry := l.touchEntry(key, now)
	entry.requests.removeExpired(now.Unix(), duration)
	if entry.requests.length+l.reservations[key] >= maxRequests {
		return nil
	}
	l.reservations[key]++
	return &Reservation{limiter: l, key: key}
}

// Complete 释放预留名额；success 为 true 时该请求计入滑动窗口。可安全重复调用。
func (r *Reservation) Complete(success bool) {
	if r == nil {
		return
	}
	r.once.Do(func() {
		l := r.limiter
		l.mutex.Lock()
		defer l.mutex.Unlock()
		l.reservations[r.key]--
		if l.reservations[r.key] == 0 {
			delete(l.reservations, r.key)
		}
		if success {
			now := time.Now()
			l.touchEntry(r.key, now).requests.append(now.Unix())
		}
	})
}
