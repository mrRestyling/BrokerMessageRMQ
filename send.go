package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go" // Делаем удобное имя для импорта в нашем коде
)

func main() {

	// Создаем подключение к RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	defer func() {
		_ = conn.Close() // Закрываем подключение в случае удачной попытки
	}()

	// Соединение является абстракцией над socket,
	// служит для согласования версии протокола между сервером и клиентом,
	// отвечает за аутентификацию и другие важные вещи.

	// Создаем канал, в котором находится большая часть логики для выполнения задач.

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	}()

	// Нужно объявить QUEUE для публикации сообщений.

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	log.Printf("Sent %s\n", body)
}
