// common/models/message.go
package models

type Message struct {
    Type    string      `json:"type"`
    Payload interface{} `json:"payload"`
}
