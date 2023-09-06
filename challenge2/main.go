package main

import (
	"challenge2/brasilapi"
	"challenge2/viacep"
	"os"
	"time"
)

func main() {

	cep := os.Args[1]

	channel1 := make(chan string)
	channel2 := make(chan string)

	brasilapiService := brasilapi.NewBrasilApi(&channel1)
	viacepService := viacep.NewViaCep(&channel2)

	go brasilapiService.GetCEP(cep)
	go viacepService.GetCEP(cep)

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
