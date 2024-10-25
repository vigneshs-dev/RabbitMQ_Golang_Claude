// payment-service/handlers/payment_handler.go
package handlers

import (
    "encoding/json"
    "log"

    "your-project/common/models"
)

func HandlePayment(message []byte) error {
    var msg models.Message
    if err := json.Unmarshal(message, &msg); err != nil {
        return err
    }

    if msg.Type == "ORDER_CREATED" {
        log.Printf("Processing payment for order: %v", msg.Payload)
        // Add payment processing logic here
    }

    return nil
}