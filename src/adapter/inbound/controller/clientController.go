package controller

import (
	"fmt"
)

type ClientController struct {
}

func NewClientController() *ClientController {
	return &ClientController{}
}

func (m *ClientController) PostMessage() {

	fmt.Println("handler")

}
