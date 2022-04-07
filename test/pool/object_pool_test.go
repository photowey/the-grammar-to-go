package poolt

import (
    "testing"
    "time"

    poolt "photowey.com/the-grammar-to-go/grammar/pool"
)

func TestObjectPool(t *testing.T) {
    bufferedSize := 10
    pool := poolt.NewObjectPool(10)
    for i := 0; i < bufferedSize; i++ {
        if object, err := pool.GetObject(time.Second * 1); err != nil {
            t.Error(err)
        } else {
            if err := pool.ReleaseObject(object); err != nil {
                t.Error(err)
            }
        }
    }

}

func TestObjectPoolAndRelease(t *testing.T) {
    bufferedSize := 10
    times := bufferedSize + 1
    pool := poolt.NewObjectPool(10)
    for i := 0; i < times; i++ {
        if object, err := pool.GetObject(time.Second * 1); err != nil {
            t.Error(err)
        } else {
            t.Logf("%T\n", object)
            if err := pool.ReleaseObject(object); err != nil {
                t.Error(err)
            }
        }
    }

    t.Log("--- Done ---")
}

func TestObjectPoolNotRelease(t *testing.T) {
    bufferedSize := 10
    times := bufferedSize + 1
    pool := poolt.NewObjectPool(10)
    for i := 0; i < times; i++ {
        // get object timeout.
        if object, err := pool.GetObject(time.Second * 1); err != nil {
            t.Error(err)
        } else {
            t.Logf("%T\n", object)

        }
    }

    t.Log("--- Done ---")
}

func TestObjectPoolReleaseMulti(t *testing.T) {
    bufferedSize := 10
    times := bufferedSize + 1
    pool := poolt.NewObjectPool(10)

    // 尝试 - 多释放
    if err := pool.ReleaseObject(&poolt.ReuseableObject{}); err != nil {
        // buffered channel overflow!
        t.Error(err)
    }

    for i := 0; i < times; i++ {
        if object, err := pool.GetObject(time.Second * 1); err != nil {
            t.Error(err)
        } else {
            t.Logf("%T\n", object)
            if err := pool.ReleaseObject(object); err != nil {
                t.Error(err)
            }
        }
    }

    t.Log("--- Done ---")
}
