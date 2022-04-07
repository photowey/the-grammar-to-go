package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func square(op int) int {
	return op * op
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		rvt := square(inputs[i])
		if rvt != expected[i] {
			t.Errorf("input is:%d,the expected is:%d, the actual is:%d", inputs[i], expected[i], rvt)
		}
	}
}

func TestSquareAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		rvt := square(inputs[i])
		assert.Equal(t, expected[i], rvt)
	}
}

func TestErrorInCode(t *testing.T) {
	t.Log("Start")
	// Error 后续代码 还会被执行
	t.Error("Error")
	t.Log("End")
}

func TestFailInCode(t *testing.T) {
	t.Log("Start")

	// Fatal 后续代码 不会被执行
	t.Fatal("Fatal")
	t.Log("End")
}
