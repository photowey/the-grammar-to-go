package structt

import (
	"fmt"
)

// Employee ...
type Employee struct {
	Id   string
	Name string
	Age  int
}

// String1 第⼀种定义⽅式在实例对应⽅法被调⽤时,实例的成员会进⾏值复制(内存拷⻉)
func (e Employee) String1() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// String 通常情况下为了避免内存拷⻉我们使⽤第⼆种定义⽅式
func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}
