package controller

import (
	"fmt"
	"net/http"
)

type ClientController struct {
}

func NewClientController() *ClientController {
	return &ClientController{}
}

func (m *ClientController) PostMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("handler")

}
