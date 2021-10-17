package groutinet

import (
	"runtime"
	"testing"

	groutinet "photowey.com/the-grammar-to-go/grammar/groutine"
)

func TestHandleGroutine(t *testing.T) {
	groutinet.HandleGroutine(5)
}

func TestHandleGroutineError(t *testing.T) {
	groutinet.HandleGroutineError(5)
}

// TestNormalChannel 返回第一个任务 - 采用通用 channel
func TestNormalChannel(t *testing.T) {
	// Before: the goroutine numbers=2
	t.Logf("Before: the goroutine numbers=%d", runtime.NumGoroutine())
	rvt := groutinet.NormalChannel(10)
	t.Log(rvt)
	// After: the goroutine numbers=11
	// -> 有很多 Goroutine 还处于阻塞状态
	// 协程泄漏
	t.Logf("After: the goroutine numbers=%d", runtime.NumGoroutine())
}

// TestBufferedChannelAll 返回所有任务 - 采用通用 buffered-channel
func TestBufferedChannelAll(t *testing.T) {
	// Before: the goroutine numbers=2
	t.Logf("Before: the goroutine numbers=%d", runtime.NumGoroutine())
	rvt := groutinet.BufferedChannelAll(10)
	t.Log(rvt)
	// After: the goroutine numbers=2
	// -> Goroutine 已经被释放
	t.Logf("After: the goroutine numbers=%d", runtime.NumGoroutine())
}
