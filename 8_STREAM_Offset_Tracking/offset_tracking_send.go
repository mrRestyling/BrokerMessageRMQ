package main

import (
	"fmt"
	"os"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func main() {

	// Создаем среду
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest"))
	CheckErrSend(err)

	// Объявляем поток
	streamName := "stream-offset-tracking-go"
	err = env.DeclareStream(streamName,
		&stream.StreamOptions{
			MaxLengthBytes: stream.ByteCapacity{}.GB(2),
		},
	)
	CheckErrSend(err)

	// Создаем продьюсьера
	producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
	CheckErrSend(err)

	messageCount := 100
	chPublishConfirm := producer.NotifyPublishConfirmation()
	ch := make(chan bool)
	handlePublishConfirm(chPublishConfirm, messageCount, ch)

	fmt.Printf("Publishing %d messages...\n", messageCount)
	for i := 0; i < messageCount; i++ {
		var body string
		if i == messageCount-1 {
			body = "marker"
		} else {
			body = "hello"
		}
		err = producer.Send(amqp.NewMessage([]byte(body)))
		CheckErrSend(err)
	}
	_ = <-ch
	fmt.Println("Messages confirmed.")

	err = producer.Close()
	CheckErrSend(err)
}

// Обработчик подтверждений
func handlePublishConfirm(confirms stream.ChannelPublishConfirm, messageCount int, ch chan bool) {
	go func() {
		confirmedCount := 0
		for confirmed := range confirms {
			for _, msg := range confirmed {
				if msg.IsConfirmed() {
					confirmedCount++
					if confirmedCount == messageCount {
						ch <- true
					}
				}
			}
		}
	}()
}

func CheckErrSend(err error) {
	if err != nil {
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}
