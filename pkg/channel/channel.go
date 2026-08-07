package channel

import (
	"context"
	"time"
)

// TrySend 尝试向 channel 发送一个值而不阻塞。
// 返回值：
//   - sent=true：发送成功
//   - sent=false：channel 已满或已关闭
//
// 注意：当 channel 已关闭时，发送会触发 panic，本函数会 recover 并返回 false，
// 调用方无法仅凭返回值区分"channel 满"与"channel 已关闭"。
// 若需区分二者，请使用 SafeSend。
func TrySend[T any](ch chan T, value T) (sent bool) {
	defer func() {
		if recover() != nil {
			sent = false
		}
	}()
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

func SafeSend[T any](ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}

// SafeSendTimeout 在指定超时时间内尝试向 channel 发送一个值。
// 优先使用 time.Timer 显式停止计时器，避免 time.After 在发送成功时造成的计时器资源延迟释放。
// 当 channel 关闭与超时同时就绪时，通过 select 语义仍可能任选其一，
// 故关闭场景下 closed=true 的判定由 defer 中的 recover 兜底保证。
func SafeSendTimeout[T any](ch chan T, value T, timeout time.Duration) (sent bool, closed bool) {
	defer func() {
		if recover() != nil {
			sent = false
			closed = true
		}
	}()
	t := time.NewTimer(timeout)
	defer t.Stop()
	select {
	case ch <- value:
		return true, false
	case <-t.C:
		return false, false
	}
}

func SafeSendContext[T any](ch chan T, value T, ctx context.Context) (sent bool, closed bool) {
	defer func() {
		if recover() != nil {
			sent = false
			closed = true
		}
	}()
	select {
	case ch <- value:
		return true, false
	case <-ctx.Done():
		return false, false
	}
}
