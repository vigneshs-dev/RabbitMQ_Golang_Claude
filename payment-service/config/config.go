// payment-service/config/config.go
package config

type Config struct {
    RabbitMQURL string
    OrderQueue  string
}

func NewConfig() *Config {
    return &Config{
        RabbitMQURL: "amqp://guest:guest@localhost:5672/",
        OrderQueue:  "orders",
    }
}