package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
)

var (
	Topic           = ""
	LotteryProducer sarama.SyncProducer // 生产消息
)

func InitKafka() {
	var err error
	address := []string{
		"ip1:port1",
		"ip2:port2",
		"ip3:port3",
		"ip4:port4",
	}
	Topic = "your topic"
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	kafkaConfig.Producer.Retry.Max = 3                    // Retry up to 10 times to produce the message
	kafkaConfig.Producer.Return.Successes = true
	LotteryProducer, err = sarama.NewSyncProducer(address, kafkaConfig)
	if err != nil {
		fmt.Println("err", err)
	}
}

// lottery send 发送kafka数据
func Send(recordId int64, data string) {
	recordIdStr := strconv.FormatInt(recordId, 10)
	//jsonData, _ := json.Marshal(data)
	//jsonData := data
	msgKafka := &sarama.ProducerMessage{
		Topic: Topic,
		Key:   sarama.StringEncoder(recordIdStr),
		Value: sarama.StringEncoder(data),
	}
	partition, offset, err := LotteryProducer.SendMessage(msgKafka)
	fmt.Println(partition, offset, err)
}

// 发送 kafka 消息，发送抽奖次数
func SendData() {
	Send(1234567, `data`)
}
