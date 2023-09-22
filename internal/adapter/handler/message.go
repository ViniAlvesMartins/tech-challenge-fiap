package handler

import (
	"fmt"
)

type MessageHandler struct {
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) PostMessage() {

	fmt.Println("handler")

}
