package rabbitmq

import (
	"log"
	"sync"
)

var (
	done chan bool
	once sync.Once
)

func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	msgs, err := channel.Consume(
		qName,
		cName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	done = make(chan bool)

	// 循环读取 channel 的数据
	go func() {
		for d := range msgs {
			processErr := callback(d.Body)
			if processErr {
				// TODO: 将错误写入错误队列
			}
		}
	}()

	// 接收 done 信号 阻塞
	<-done

	if err := channel.Close(); err != nil {
		log.Println("channel.Close err", err)
	}
}

func StopConsume() {
	once.Do(func() {
		close(done)
	})
}
