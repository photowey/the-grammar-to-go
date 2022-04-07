package methodset

import (
    `fmt`
    `reflect`
)

func DumpMethodSet(i interface{}) (n int) {
    dynamicType := reflect.TypeOf(i)
    if dynamicType == nil {
        fmt.Printf("there is no dynamic type\n")
        return
    }
    n = dynamicType.NumMethod()
    if n == 0 {
        fmt.Printf("%s's method set is empty!\n", dynamicType)
        return
    }
    fmt.Printf("%s's method set:\n", dynamicType)
    for j := 0; j < n; j++ {
        fmt.Println("-", dynamicType.Method(j).Name)
    }
    fmt.Printf("\n")

    return
}
