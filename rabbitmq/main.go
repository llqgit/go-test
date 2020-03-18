package main

import (
	"fmt"
	"github.com/assembla/cony"
	"github.com/streadway/amqp"
	"net"
	"time"
)

var (
	Cli *cony.Client
	Cns *cony.Consumer
)

type QMessage struct {
	Message string `json:"message"`
}

func init() {
	giftMq()
}

func giftMq() {
	//连接到rabbitmq

	//rabbitmq.url=amqp://admin:7gNPQFqVp7w20kYZ@10.110.40.136:5672
	//rabbitmq.exchange=gift_topic
	//rabbitmq.queue=gift
	url := "amqp://admin:7gNPQFqVp7w20kYZ@10.110.40.136:5672"
	queName := "gift_test_llq_1"
	excName := "gift_topic"
	config := amqp.Config{
		Heartbeat: time.Second * 6,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 2*time.Second)
		},
	}
	Cli = cony.NewClient(
		cony.URL(url),
		cony.Backoff(cony.DefaultBackoff),
		cony.Config(config),
	)

	que := &cony.Queue{
		Name:       queName, // name
		Durable:    true,    // durable
		AutoDelete: false,   // delete when unused
		Exclusive:  false,   // exclusive
		Args:       nil,
	}
	exc := cony.Exchange{
		Name:    excName,
		Kind:    "topic",
		Durable: true,
	}

	bind := cony.Binding{
		Queue:    que,
		Exchange: exc,
		Key:      "*",
	}
	Cli.Declare([]cony.Declaration{
		cony.DeclareExchange(exc),
		cony.DeclareQueue(que),
		cony.DeclareBinding(bind),
	})
	// Declare and register a consumer
	Cns = cony.NewConsumer(
		que,
		cony.Qos(5000),
	)
	Cli.Consume(Cns)

	fmt.Println("load mq")
}

func main() {
	fmt.Println("running......")
	for Cli.Loop() {
		select {
		case d := <-Cns.Deliveries():
			fmt.Println("get gift", d)
			_ = d.Ack(false)
			//go actions.DealWith(d)

		case err := <-Cns.Errors():
			//raven.CaptureError(err, nil)
			fmt.Println(err)
		case err := <-Cli.Errors():
			//raven.CaptureError(err, nil)
			fmt.Println(err)
		}
	}
}
