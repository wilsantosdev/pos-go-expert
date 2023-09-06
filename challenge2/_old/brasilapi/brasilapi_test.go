package brasilapi

import "testing"

func TestMakeUrl(t *testing.T) {
	var brasilAPIChannel = make(chan string)
	brasilAPI := NewBrasilAPI(brasilAPIChannel)

	url := brasilAPI.makeUrl("23017-335")

	if url != "https://brasilapi.com.br/api/cep/v1/23017-335" {
		t.Errorf("URL is not valid")
	}
}

func TestGetCEP(t *testing.T) {
	cep_expected := `{"cep":"23017335","state":"RJ","city":"Rio de Janeiro","neighborhood":"Campo Grande","street":"Rua Valdir Pequeno de Melo","service":"correios"}`

	var brasilAPIChannel = make(chan string)
	brasilAPI := NewBrasilAPI(brasilAPIChannel)

	cep, err := brasilAPI.GetCEP("23017-335")

	if err != nil {
		t.Errorf("Error on get cep")
	}

	if cep != cep_expected {
		t.Errorf("Cep is not valid")
	}

}
