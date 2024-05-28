package rabbitmq

import (
	"cloud_storage/configurator"
	"github.com/streadway/amqp"
	"log"
)

var notifyClose chan *amqp.Error

func init() {
	if initChannel() {
		channel.NotifyClose(notifyClose)
	}

	// 断线重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel()
			}
		}
	}()
}

func initChannel() bool {
	if channel != nil {
		return true
	}
	cgf := configurator.GetRabbitMQConfig()
	conn, err := amqp.Dial(cgf.GetAttr("Url"))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel() {
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plian",
			Body:        msg}) {
		return true
	}
	return false
}
