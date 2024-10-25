// order-service/handlers/order_handler.go
package handlers

import (
    "encoding/json"
    "log"

    "your-project/common/models"
    "your-project/common/rabbitmq"
    orderModels "your-project/order-service/models"
)

type OrderHandler struct {
    rabbitmq *rabbitmq.RabbitMQ
}

func NewOrderHandler(rmq *rabbitmq.RabbitMQ) *OrderHandler {
    return &OrderHandler{
        rabbitmq: rmq,
    }
}

func (h *OrderHandler) CreateOrder(order orderModels.Order) error {
    message := models.Message{
        Type:    "ORDER_CREATED",
        Payload: order,
    }

    return h.rabbitmq.PublishMessage("orders", message)
}