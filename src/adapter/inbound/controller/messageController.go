package controller

import (
	"fmt"
)

type MessageController struct {
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

func (m *MessageController) PostMessage() {

	fmt.Println("handler")

}
