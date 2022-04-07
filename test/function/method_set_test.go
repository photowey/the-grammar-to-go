package function

import (
    `testing`

    `github.com/stretchr/testify/assert`
    `photowey.com/the-grammar-to-go/grammar/methodset`
)

type T struct{}

func (T) M1() {}
func (T) M2() {}

type S T
type S1 = T

func TestDumpMethodSet(t *testing.T) {

    asserts := assert.New(t)

    s := S{}
    s1 := S1{}
    msc1 := methodset.DumpMethodSet(s)
    asserts.Equal(0, msc1)

    msc2 := methodset.DumpMethodSet(s1)
    asserts.Equal(2, msc2)
}
