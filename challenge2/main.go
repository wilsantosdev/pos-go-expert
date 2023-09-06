package main

import (
	"challenge2/brasilapi"
	"challenge2/viacep"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	brasilapiService := brasilapi.NewBrasilApi(&channel1)
	viacepService := viacep.NewViaCep(&channel2)

	go brasilapiService.GetCEP("01001000")
	go viacepService.GetCEP("01001000")

	select {
	case response := <-channel1:
		println("BRASILAPI")
		println("=====================================")
		println(response)
	case response := <-channel2:
		println("VIACEP")
		println("=====================================")
		println(response)
	case <-time.After(1 * time.Second):
		println("TIMEOUT")
	}

}
