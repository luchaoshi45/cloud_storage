package rabbitmq

import "github.com/streadway/amqp"

// TransferData : 将要写到rabbitmq的数据的结构体
type TransferData struct {
	FileHash     string
	CurLocation  string
	DestLocation string
}

var conn *amqp.Connection
var channel *amqp.Channel
