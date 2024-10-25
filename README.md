# Go Microservices with RabbitMQ

This project demonstrates a microservices architecture using Go and RabbitMQ for inter-service communication.

## Services

1. Order Service: Handles order creation and management
2. Payment Service: Processes payments for orders

## Prerequisites

- Go 1.21 or higher
- Docker (for RabbitMQ)
- Git

## Setup

1. Clone the repository:
\`\`\`bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
\`\`\`

2. Start RabbitMQ:
\`\`\`bash
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
\`\`\`

3. Install dependencies:
\`\`\`bash
go mod tidy
\`\`\`

## Running the Services

1. Start Order Service:
\`\`\`bash
cd order-service
go run main.go
\`\`\`

2. Start Payment Service:
\`\`\`bash
cd payment-service
go run main.go
\`\`\`

## Project Structure

\`\`\`
├── order-service/
│   ├── main.go
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   └── order.go
│   └── handlers/
│       └── order_handler.go
├── payment-service/
│   ├── main.go
│   ├── config/
│   │   └── config.go
│   └── handlers/
│       └── payment_handler.go
├── common/
│   ├── rabbitmq/
│   │   └── rabbitmq.go
│   └── models/
│       └── message.go
└── go.mod
\`\`\`

## License

MIT License