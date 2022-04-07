package selectt

import (
    "testing"
    "time"

    selectt "photowey.com/the-grammar-to-go/grammar/select"
)

func TestSelect(t *testing.T) {
    select {
    case rvt := <-selectt.AsyncService():
        t.Log("handle AsyncService result is:", rvt)
    case <-time.After(time.Millisecond * 600):
        t.Error("handle AsyncService() time out~")
    }
}
