// for.go

package fori

import (
    "fmt"
)

// Go 语⾔仅⽀持循环关键字 for

const (
    // TimesThreshold 循环次数
    TimesThreshold int = 2
)

var (
    // CONTEXT 不指定元素个数
    CONTEXT = [...]int{1, 2, 3, 4, 5}
)

// ForI {@code for-i } 循环
func ForI() {
    for i := 0; i < 3; i++ {
        fmt.Printf("for-i index:%d\n", i)
    }
    fmt.Println()
}

// ForRange {@code for-range } 循环
func ForRange() {
    var greeting = "Hello go!"
    for _, elem := range greeting {
        fmt.Printf("%c", elem)
    }
}

// ForWhile for 无限循环
func ForWhile() {
    n := 0
    for {
        if n > TimesThreshold {
            break
        }
        n++
        fmt.Printf("--- exec for while-true n==%d---\n", n)
    }
}

// TravelArray 遍历数组
func TravelArray() (sum int) {
    for idx /*索引*/, elem /*元素*/ := range CONTEXT {
        fmt.Printf("the index:%d,value is:[%d]", idx, elem)
        sum += elem
    }
    return
}
