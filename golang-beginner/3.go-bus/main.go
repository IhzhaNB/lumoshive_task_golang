package main

import (
	"go-bus/dto"
	"go-bus/handler"
)

func main() {
	req := dto.NewRequest("Sidik", "Jakbar")

	handler.ProcessTicketRequest(req)
}
