package handler

import (
	"fmt"
	"go-bus/dto"
	"go-bus/service"
)

func ProcessTicketRequest(req *dto.Request) {
	response := service.CheckPrice(req)

	fmt.Println("=== Harga Tiket ===")

	// Cek Found yg diakses
	if !response.Found {
		fmt.Printf("Maaf, destinasi '%s' tidak ditemukan dalam data.\n", response.Destination)
		return
	}

	fmt.Printf("Penumpang	: %s\n", response.PassengerName)
	fmt.Printf("Tujuan		: %s\n", response.Destination)
	fmt.Printf("Harga		: Rp %.2f\n", response.Price)
}
