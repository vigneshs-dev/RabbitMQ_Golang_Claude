// payment-service/main.go
package main

import (
    "log"

    "RabbitMQ_Golang_Claude/common/rabbitmq"
    "RabbitMQ_Golang_Claude/payment-service/config"
    "RabbitMQ_Golang_Claude/payment-service/handlers"
)

func main() {
    cfg := config.NewConfig()

    rmq, err := rabbitmq.NewRabbitMQ(cfg.RabbitMQURL)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer rmq.Close()

    if err := rmq.DeclareQueue(cfg.OrderQueue); err != nil {
        log.Fatalf("Failed to declare queue: %v", err)
    }

    log.Println("Payment service is running...")
    rmq.ConsumeMessages(cfg.OrderQueue, handlers.HandlePayment)
}