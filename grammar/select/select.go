package selectt

import (
    "fmt"
    "time"
)

func syncService() string {
    time.Sleep(time.Millisecond * 500)

    return "syncService() handle Done"
}

func AsyncService() chan string {
    rvtCh := make(chan string, 1)
    go func() {
        rvt := syncService()
        fmt.Println("the syncService() returned result...")
        rvtCh <- rvt
        fmt.Println("the syncService() exit after rvtCh <-...")
    }()

    return rvtCh
}
