// function_test.go

package function

import (
    functiont "photowey.com/the-grammar-to-go/grammar/function"
    "testing"
)

// TestVariadicParameterFunctionSum 测试-可变参数函数求和
func TestVariadicParameterFunctionSum(t *testing.T) {
    sum := functiont.VariadicParameterFunctionSum(1, 2, 3, 4)
    t.Logf("the test variadic parameter function sum result is: [%d]", sum)
}

// TestDeferFunction 测试 {@code defer-function}
func TestDeferFunction(t *testing.T) {
    defer func() {
        t.Log("Clear resources")
    }()
    t.Log("Started")
    // defer 仍会执⾏
    panic("Fatal error")
}
