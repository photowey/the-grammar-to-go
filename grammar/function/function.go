// function.go

package function

// 1.可以有多个返回值
// 2.所有参数都是值传递：slice,map,channel 会有传引⽤的错觉
// 3.函数可以作为变量的值
// 4.函数可以作为参数和返回值

// VariadicParameterFunctionSum 可变参数-函数-求和
func VariadicParameterFunctionSum(ops ...int) (sum int) {
    sum = 0
    for _, op := range ops {
        sum += op
    }
    return
}
