package rabbitmq

import "nakashima/rabbitmq/service"

// NewRabbitMq 创建mq对象
// uri mq连接字符串
// queue 队列名
// amqp://user:pwd@host:port/vhost
func NewRabbitMq(uri string, queue string, exChange string) *service.RabbitMQ {
	return service.NewRabbitMQSimple(uri, queue, exChange)
}
