package structt

import (
	structt "photowey.com/the-grammar-to-go/grammar/struct"
	"testing"
)

// TestStructEmployee 测试-结构体-{@code Employee}
func TestStructEmployee(t *testing.T) {
	// 推荐采用 e2 的方式实例化-结构体
	// e1 := structt.Employee{"8848", "photowey", 18}
	// t.Log(e1)

	e2 := structt.Employee{Id: "9527", Name: "photowey", Age: 20}
	t.Log(e2)
	t.Logf("the type of e2 is:%T", e2)
	// 注意这⾥返回的引⽤/指针，相当于 e2 := &Employee{}
	e3 := new(structt.Employee)
	e3.Id = "1024"
	e3.Name = "photowey"
	e3.Age = 22

	t.Log(e3)
	t.Logf("the type of e3 is:%T", e3)

	e4 := &structt.Employee{Id: "9527", Name: "photowey", Age: 20}
	// 调用 e4.String() 函数
	t.Log("the value of e4 is:", e4)
	t.Log("the value of *e4 is:", *e4)
	t.Logf("the type of e4 is:%T", e4)
}
