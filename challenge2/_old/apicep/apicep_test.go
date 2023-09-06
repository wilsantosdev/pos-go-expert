package apicep

import "testing"

func TestMakeUrl(t *testing.T) {
	var apiCEPChannel = make(chan string)
	apiCep := NewApiCep(apiCEPChannel)

	url := apiCep.makeUrl("23017-335")

	if url != "https://cdn.apicep.com/file/apicep/23017-335.json" {
		t.Errorf("URL is not valid")
	}
}

func TestGetCEP(t *testing.T) {
	cep_expected := `{"code":"23017-335","state":"RJ","city":"Rio de Janeiro","district":"Campo Grande","address":"Rua Valdir Pequeno de Melo","status":200,"ok":true,"statusText":"ok"}`

	var apiCEPChannel = make(chan string)
	apiCep := NewApiCep(apiCEPChannel)

	cep, err := apiCep.GetCEP("23017-335")

	if err != nil {
		t.Errorf("Error on get cep")
	}

	if cep != cep_expected {
		t.Errorf("Cep is not valid")
	}
}
