package poolt

import (
    "errors"
    "fmt"
    "time"
)

// ReuseableObject reuseable object
type ReuseableObject struct {
    Id   int
    Name string
    Age  int
}

type ObjectPool struct {
    // 用于-缓存可复用对象
    bufferedChannel chan *ReuseableObject
}

// NewObjectPool create a buffered-object-pool
func NewObjectPool(bufferedSize int) *ObjectPool {
    objectPool := ObjectPool{}
    objectPool.bufferedChannel = make(chan *ReuseableObject, bufferedSize)
    for i := 0; i < bufferedSize; i++ {
        objectPool.bufferedChannel <- &ReuseableObject{Id: 9527 + i, Name: fmt.Sprintf("photowey %d", i), Age: 18}
    }

    return &objectPool
}

// GetObject get object from buffered-channel object-pool
func (pool *ObjectPool) GetObject(timeout time.Duration) (*ReuseableObject, error) {
    select {
    case rvt := <-pool.bufferedChannel:
        return rvt, nil
    case <-time.After(timeout):
        // 超时控制
        return nil, errors.New("get object timeout.")
    }
}

// ReleaseObject release the object to object-pool
func (pool *ObjectPool) ReleaseObject(object *ReuseableObject) error {
    select {
    case pool.bufferedChannel <- object:
        return nil
    default:
        return errors.New("buffered channel overflow!")
    }
}
