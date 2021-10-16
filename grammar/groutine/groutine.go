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
