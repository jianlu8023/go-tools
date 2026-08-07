package channel

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTrySend(t *testing.T) {
	t.Run("send success", func(t *testing.T) {
		ch := make(chan int, 1)
		result := TrySend(ch, 1)
		if !result {
			t.Errorf("expected true, got false")
		}
		v := <-ch
		if v != 1 {
			t.Errorf("expected 1, got %d", v)
		}
	})

	t.Run("send fail with full channel", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 1
		result := TrySend(ch, 2)
		if result {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("concurrent send", func(t *testing.T) {
		ch := make(chan int, 10)
		var wg sync.WaitGroup
		var sent int32

		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				if TrySend(ch, val) {
					atomic.AddInt32(&sent, 1)
				}
			}(i)
		}

		wg.Wait()
		if sent != 10 {
			t.Errorf("expected 10 sent, got %d", sent)
		}
	})
}

func TestSafeSend(t *testing.T) {
	t.Run("send success", func(t *testing.T) {
		ch := make(chan int, 1)
		closed := SafeSend(ch, 1)
		if closed {
			t.Errorf("expected false, got true")
		}
		v := <-ch
		if v != 1 {
			t.Errorf("expected 1, got %d", v)
		}
	})

	t.Run("send fail with closed channel", func(t *testing.T) {
		ch := make(chan int)
		close(ch)
		closed := SafeSend(ch, 1)
		if !closed {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("concurrent send", func(t *testing.T) {
		ch := make(chan int, 20)
		var wg sync.WaitGroup
		var closedCount int32

		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				if SafeSend(ch, val) {
					atomic.AddInt32(&closedCount, 1)
				}
			}(i)
		}

		wg.Wait()
		if closedCount != 0 {
			t.Errorf("expected 0 closed, got %d", closedCount)
		}
	})

	t.Run("concurrent send to closed channel", func(t *testing.T) {
		ch := make(chan int)
		close(ch)
		var wg sync.WaitGroup
		var closedCount int32

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if SafeSend(ch, i) {
					atomic.AddInt32(&closedCount, 1)
				}
			}()
		}

		wg.Wait()
		if closedCount != 10 {
			t.Errorf("expected 10 closed, got %d", closedCount)
		}
	})
}

func TestSafeSendTimeout(t *testing.T) {
	t.Run("send success", func(t *testing.T) {
		ch := make(chan int, 1)
		sent, closed := SafeSendTimeout(ch, 1, 100*time.Millisecond)
		if !sent {
			t.Errorf("expected sent=true, got false")
		}
		if closed {
			t.Errorf("expected closed=false, got true")
		}
		v := <-ch
		if v != 1 {
			t.Errorf("expected 1, got %d", v)
		}
	})

	t.Run("send timeout", func(t *testing.T) {
		ch := make(chan int)
		sent, closed := SafeSendTimeout(ch, 1, 50*time.Millisecond)
		if sent {
			t.Errorf("expected sent=false, got true")
		}
		if closed {
			t.Errorf("expected closed=false, got true")
		}
	})

	t.Run("send fail with closed channel", func(t *testing.T) {
		ch := make(chan int)
		close(ch)
		sent, closed := SafeSendTimeout(ch, 1, 100*time.Millisecond)
		if sent {
			t.Errorf("expected sent=false, got true")
		}
		if !closed {
			t.Errorf("expected closed=true, got false")
		}
	})

	t.Run("concurrent send with timeout", func(t *testing.T) {
		ch := make(chan int, 5)
		var wg sync.WaitGroup
		var sentCount int32
		var timeoutCount int32

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				sent, _ := SafeSendTimeout(ch, val, 50*time.Millisecond)
				if sent {
					atomic.AddInt32(&sentCount, 1)
				} else {
					atomic.AddInt32(&timeoutCount, 1)
				}
			}(i)
		}

		wg.Wait()
		if sentCount+timeoutCount != 10 {
			t.Errorf("expected total 10, got sent=%d, timeout=%d", sentCount, timeoutCount)
		}
	})
}

func TestSafeSendContext(t *testing.T) {
	t.Run("send success", func(t *testing.T) {
		ch := make(chan int, 1)
		ctx := context.Background()
		sent, closed := SafeSendContext(ch, 1, ctx)
		if !sent {
			t.Errorf("expected sent=true, got false")
		}
		if closed {
			t.Errorf("expected closed=false, got true")
		}
		v := <-ch
		if v != 1 {
			t.Errorf("expected 1, got %d", v)
		}
	})

	t.Run("send cancelled", func(t *testing.T) {
		ch := make(chan int)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		sent, closed := SafeSendContext(ch, 1, ctx)
		if sent {
			t.Errorf("expected sent=false, got true")
		}
		if closed {
			t.Errorf("expected closed=false, got true")
		}
	})

	t.Run("send fail with closed channel", func(t *testing.T) {
		ch := make(chan int)
		close(ch)
		ctx := context.Background()
		sent, closed := SafeSendContext(ch, 1, ctx)
		if sent {
			t.Errorf("expected sent=false, got true")
		}
		if !closed {
			t.Errorf("expected closed=true, got false")
		}
	})

	t.Run("send with timeout context", func(t *testing.T) {
		ch := make(chan int)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		sent, closed := SafeSendContext(ch, 1, ctx)
		if sent {
			t.Errorf("expected sent=false, got true")
		}
		if closed {
			t.Errorf("expected closed=false, got true")
		}
	})
}
