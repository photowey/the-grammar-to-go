package channelt

import (
	"sync"
	"testing"

	channelt "photowey.com/the-grammar-to-go/grammar/channel"
)

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	// 单生产者 - 生产
	channelt.DataProducer(ch, 10, &wg)

	// 多消费者 - 消费
	wg.Add(1)
	channelt.DataConsumer(ch, 1111, &wg)
	wg.Add(1)
	channelt.DataConsumer(ch, 2222, &wg)
	wg.Add(1)
	channelt.DataConsumer(ch, 3333, &wg)

	wg.Wait()
}
