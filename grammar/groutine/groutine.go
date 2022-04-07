package groutinet

import (
    "fmt"
    "time"
)

func HandleGroutine(times int) {
    for i := 0; i < times; i++ {
        go func(index int) {
            fmt.Printf("receive the index is:%d\n", index)
        }(i)
    }
    time.Sleep(time.Millisecond * 50)
}

func HandleGroutineError(times int) {
    for i := 0; i < times; i++ {
        go func() {
            // loop variable i captured by func literal loopclosure
            fmt.Printf("receive the index is:%d\n", i)
        }()
    }
    time.Sleep(time.Millisecond * 50)
}

func runTask(id int) string {
    time.Sleep(time.Millisecond * 10)
    return fmt.Sprintf("the result is from id:%d", id)
}

// NormalChannel use normal channel
func NormalChannel(times int) string {
    ch := make(chan string)
    for i := 0; i < times; i++ {
        go func(index int) {
            rvt := runTask(index)
            ch <- rvt
        }(i)
    }

    return <-ch
}

// BufferedChannel use buffered-channel
func BufferedChannel(times int) string {
    ch := make(chan string, times)
    for i := 0; i < times; i++ {
        go func(index int) {
            rvt := runTask(index)
            ch <- rvt
        }(i)
    }

    return <-ch
}

// BufferedChannelAll use buffered-channel
func BufferedChannelAll(times int) (finalResult string) {
    ch := make(chan string, times)
    for i := 0; i < times; i++ {
        go func(index int) {
            rvt := runTask(index)
            ch <- rvt
        }(i)
    }

    for j := 0; j < times; j++ {
        finalResult += <-ch + "\n"
    }

    return finalResult
}
