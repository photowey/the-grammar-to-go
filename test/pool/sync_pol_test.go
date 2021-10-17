package poolt

import (
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("create a new object by New()")
			return 101
		},
	}

	value, _ := pool.Get().(int)
	// create a new object by New()
	// the pool.Get() value is: 101
	t.Log("the pool.Get() value is:", value)

	pool.Put(10001)
	valuePut, _ := pool.Get().(int)
	// the pool.Get() valuePut is: 10001
	t.Log("the pool.Get() valuePut is:", valuePut)

	// GC 会清除 sync.Pool 里面的缓存对象
	runtime.GC()

	valueGc, _ := pool.Get().(int)
	// create a new object by New()
	// the pool.Get()after GC, valueGc is: 101
	t.Log("the pool.Get()after GC, valueGc is:", valueGc)
}

func TestSyncPoolGouroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("create a new object by New()")
			return 101
		},
	}

	pool.Put(10001)
	pool.Put(10001)
	pool.Put(10001)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
