package service

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	// 连接信息
	Mqurl string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(uri string, queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: uri}
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败！")
	return rabbitmq
}

// Destory 断开channel和connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMQSimple 简单模式step1： 1.创建简单模式下的rabbitmq实例
func NewRabbitMQSimple(uri string, queueName string, exChange string) *RabbitMQ {
	return NewRabbitMQ(uri, queueName, exChange, "")
}

// PublishSimple 简单模式step2： 2.简单模式下生产
func (r *RabbitMQ) PublishSimple(message string) {
	// 1. 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true,根据exchange类型和routekey规则，如果无法找到符合条件的队列，则会把发送的消息返回给发送者
		false,
		// 如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{ContentType: "application/json", Body: []byte(message)},
	)
}

func (r *RabbitMQ) SimpleGet() amqp.Delivery {
	msg, ok, err := r.channel.Get(r.QueueName, true)
	if !ok {
		r.failOnErr(err, "获取失败")
	}
	return msg
}
