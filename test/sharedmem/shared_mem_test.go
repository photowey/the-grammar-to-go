package sharedmem

import (
    "testing"

    "photowey.com/the-grammar-to-go/grammar/sharedmem"
)

func TestCounter(t *testing.T) {
    counter := sharedmem.Counter()
    t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
    counter := sharedmem.CounterThreadSafe()
    t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
    counter := sharedmem.CounterWaitGroup()
    t.Logf("counter = %d", counter)
}
