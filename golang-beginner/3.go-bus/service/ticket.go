package service

import (
	"encoding/json"
	"fmt"
	"go-bus/data"
	"go-bus/dto"
)

func CheckPrice(req *dto.Request) *dto.Response {
	// Deklarasi variable destinations dengan string sbg key nya, dan float64 sbg valuenya
	var destinationPrices map[string]float64

	err := json.Unmarshal([]byte(data.Destination), &destinationPrices)
	if err != nil {
		fmt.Println("Error decoding data JSON:", err)
		return &dto.Response{Found: false}
	}

	// Mencari harga di map yang sudah di-decode
	price, found := destinationPrices[req.Destination]

	return &dto.Response{
		PassengerName: req.Name,
		Destination:   req.Destination,
		Price:         price,
		Found:         found,
	}
}
