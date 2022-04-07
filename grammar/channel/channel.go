package channelt

import (
	"fmt"
	"sync"
)

func DataProductor(ch chan int, times int, wg *sync.WaitGroup) {

	go func() {
		for i := 0; i < times; i++ {
			ch <- i
		}

		// 关闭 channel
		// 通过 关闭 channel 来告诉消费者 - 我已经生产完了
		close(ch)

		// channel 关闭后 - 继续发送数据 -> 引发 panic
		// panic: send on closed channel
		// ch <- 1101

		wg.Done()
	}()

}

func DataConsumer(ch chan int, no int, wg *sync.WaitGroup) {

	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("the consumer-no:%d receive data is:%d\n", no, data)
			} else {
				break
			}
		}
		wg.Done()
	}()

}
