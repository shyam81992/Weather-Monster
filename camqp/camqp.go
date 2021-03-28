package camqp

import (
	"fmt"

	"github.com/shyam81992/Weather-Monster/config"
	"github.com/streadway/amqp"
)

//go:generate mockgen -destination=./mock/camqp.go -package=mock github.com/shyam81992/Weather-Monster/camqp CAMQPinterface,CAMQPConninterface,CAMQPChannelinterface

type CAMQP struct {
}

func (a *CAMQP) Publishmsg(msg []byte) {
	conn, err := a.Dial(config.RabbitConfig["uri"])
	if err != nil {
		fmt.Println(err, "Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err, "Failed to open a channel")
	}
	defer ch.Close()
	err = ch.Publish(
		"",                               // exchange
		config.RabbitConfig["queuename"], // routing key
		false,                            // mandatory
		false,                            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		})

	if err != nil {
		fmt.Println(err, "Failed to publish a message")
	}
}

func (a *CAMQP) Dial(url string) (CAMQPConninterface, error) {
	var ch CAMQPConn
	var err error
	ch.Connection, err = amqp.Dial(url)
	return &ch, err
}

type CAMQPinterface interface {
	Publishmsg([]byte)
}

type CAMQPConn struct {
	Connection *amqp.Connection
}

func (conn *CAMQPConn) Close() error {
	return conn.Connection.Close()
}

func (conn *CAMQPConn) Channel() (CAMQPChannelinterface, error) {
	return conn.Connection.Channel()
}

type CAMQPConninterface interface {
	Close() error
	Channel() (CAMQPChannelinterface, error)
}

type CAMQPChannelinterface interface {
	Close() error
	Publish(string, string, bool, bool, amqp.Publishing) error
}
