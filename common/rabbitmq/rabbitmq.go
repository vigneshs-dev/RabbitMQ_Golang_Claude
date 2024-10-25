// common/rabbitmq/rabbitmq.go
package rabbitmq

import (
    "encoding/json"
    "log"

    "github.com/streadway/amqp"
)

type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitMQ{
        conn:    conn,
        channel: ch,
    }, nil
}

func (r *RabbitMQ) DeclareQueue(name string) error {
    _, err := r.channel.QueueDeclare(
        name,  // name
        true,  // durable
        false, // delete when unused
        false, // exclusive
        false, // no-wait
        nil,   // arguments
    )
    return err
}

func (r *RabbitMQ) PublishMessage(queueName string, message interface{}) error {
    body, err := json.Marshal(message)
    if err != nil {
        return err
    }

    return r.channel.Publish(
        "",       // exchange
        queueName, // routing key
        false,    // mandatory
        false,    // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:       body,
        },
    )
}

func (r *RabbitMQ) ConsumeMessages(queueName string, handler func([]byte) error) {
    msgs, err := r.channel.Consume(
        queueName, // queue
        "",       // consumer
        true,     // auto-ack
        false,    // exclusive
        false,    // no-local
        false,    // no-wait
        nil,      // args
    )
    if err != nil {
        log.Printf("Failed to consume messages: %v", err)
        return
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            if err := handler(d.Body); err != nil {
                log.Printf("Error handling message: %v", err)
            }
        }
    }()

    <-forever
}

func (r *RabbitMQ) Close() {
    if r.channel != nil {
        r.channel.Close()
    }
    if r.conn != nil {
        r.conn.Close()
    }
}