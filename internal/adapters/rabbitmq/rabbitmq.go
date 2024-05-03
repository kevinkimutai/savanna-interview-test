package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewRabbitMQServer(serverurl string) *RabbitMQAdapter {
	conn, err := amqp.Dial(serverurl)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		"sms_queue", // name
		"fanout",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	return &RabbitMQAdapter{
		conn: conn,
		ch:   ch,
	}
}

func (q *RabbitMQAdapter) SendSMSQueue(order domain.Order, phoneNumber uint, customerName string) {
	//Construct Message
	msg := fmt.Sprintf("Hi! %s\n,Your order %s,totalling %s, has been successfully placed. Delivery is 2-3 days",
		customerName,
		order.OrderID,
		strconv.FormatFloat(order.TotalAmount, 'f', 1, 64),
	)

	// Create a map with message and phone number
	body := map[string]interface{}{
		"msg":         msg,
		"phonenumber": phoneNumber,
	}

	// Encode the body map as JSON
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to encode body map as JSON: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = q.ch.PublishWithContext(ctx,
		"sms_queue", // exchange
		"",          // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyJSON),
		})
	failOnError(err, "Failed to publish a message")

	log.Println(" [x][x][x][x][x] Sent SMSQueue to RabbitMQ  [x][x][x][x][x] ")
}
