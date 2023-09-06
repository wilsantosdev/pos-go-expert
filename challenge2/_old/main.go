package main

import (
	"os"
	"regexp"

	"challenge2/viacep"
)

func main() {
	cep := os.Args[1]

	if !validateCepFormat(cep) {
		println("CEP is not valid")
		return
	}

	var channel1 = make(chan string)
	// var channel2 = make(chan string)

	viacepService := viacep.NewViaCep(&channel1)
	// brasilapiService := brasilapi.NewBrasilAPI(&channel2)

	cepResponse, err := viacepService.GetCEP(cep)
	if err != nil {
		println(err.Error())
		return
	}
	println(cepResponse)
	println("from viacep")
	println("====================================")

	// cepResponse, err = brasilapiService.GetCEP(cep)
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }

	// println(cepResponse)
	// println("from brasilapi")
	// println("====================================")

	// create a goroutine for each service
	// show response on terminal
	// show error if timeout reached in 1 second

	println(cepResponse)
}

func validateCepFormat(cep string) bool {
	match, _ := regexp.MatchString(`^\d{5}-?\d{3}$`, cep)
	return match
}
