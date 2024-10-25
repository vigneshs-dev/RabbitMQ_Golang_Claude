// order-service/main.go
package main

import (
    "log"

    "RabbitMQ_Golang_Claude/common/rabbitmq"
    "RabbitMQ_Golang_Claude/order-service/config"
    "RabbitMQ_Golang_Claude/order-service/handlers"
    "RabbitMQ_Golang_Claude/order-service/models"
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

    orderHandler := handlers.NewOrderHandler(rmq)

    // Example: Create a new order
    order := models.Order{
        ID:     "123",
        UserID: "user_123",
        Amount: 99.99,
        Status: "pending",
    }

    if err := orderHandler.CreateOrder(order); err != nil {
        log.Printf("Failed to create order: %v", err)
    }

    log.Println("Order service is running...")
    select {}
}