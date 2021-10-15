package mapt

import (
	"testing"

	mapt "photowey.com/the-grammar-to-go/grammar/map"
)

// TestMapWithFuncValue 测试-函数作为 {@code map} 的 {@code value} 值
func TestMapWithFuncValue(t *testing.T) {
	m := mapt.FuncValueMap()
	t.Log(m[1](2), m[2](2), m[3](2))
}

// TestMapAsSet 测试-通过 {@code map} 实现 Set 集合
func TestMapAsSet(t *testing.T) {
	s := mapt.Set{Container: make(map[interface{}]bool)}
	s.Add(1)
	t.Logf("the set size:[%d]", s.Size())
	t.Logf("the set contains the key:[%d]?:%t", 1, s.Contains(1))
	s.Add(2)
	t.Logf("after add 2,the set size:[%d]: expect == 2", s.Size())
	s.Add(1)
	t.Logf("after add 1 again,the set size:[%d]: expect == 2", s.Size())
	s.Delete(1)
	t.Logf("after delete 1,the set size:[%d]", s.Size())
	s.Add(1)
	t.Logf("after delete and add, the set size:[%d]", s.Size())
}
