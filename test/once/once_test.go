package oncet

import (
    "sync"
    "testing"
    "unsafe"

    oncet "photowey.com/the-grammar-to-go/grammar/once"
)

func TestGetSingletonInstance(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(index int) {
            instance := oncet.GetInstance()
            // the index:9, instance &instance=824634286080
            t.Logf("the index:%d, instance &instance=%d", index, unsafe.Pointer(instance))
            wg.Done()
        }(i)
    }
    wg.Wait()
}
