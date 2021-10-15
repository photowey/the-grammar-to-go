// for_test.go

package fori

import (
	"testing"

	fort "photowey.com/the-grammar-to-go/grammar/fori"
)

// 源码⽂件以 _test 结尾：xxx_test.go
// 测试⽅法名以 Test 开头：func TestXXxx(t *testing.T) {…}

// TestForI 测试-{@code for-i} 循环
func TestForI(t *testing.T) {
	fort.ForI()
}

// TestForRange 测试-{@code for-range} 循环
func TestForRange(t *testing.T) {
	fort.ForRange()
}

// TestForWhile 测试-无限循环
func TestForWhile(t *testing.T) {
	fort.ForWhile()
}

// TestTravelArray 测试-遍历数组
func TestTravelArray(t *testing.T) {
	sum := fort.TravelArray()

	t.Log("test travel the array and sum, the sum result is:", sum)
}
