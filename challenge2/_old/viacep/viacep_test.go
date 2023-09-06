package viacep

import "testing"

func TestMakeUrl(t *testing.T) {
	var viaCEPChannel = make(chan string)
	viaCep := NewViaCep(viaCEPChannel)

	url := viaCep.makeUrl("23017-335")

	if url != "http://viacep.com.br/ws/23017-335/json/" {
		t.Errorf("URL is not valid")
	}
}

func TestGetCEP(t *testing.T) {
	cep_expected := "{\n  \"cep\": \"23017-335\",\n  \"logradouro\": \"Rua Valdir Pequeno de Melo\",\n  \"complemento\": \"\",\n  \"bairro\": \"Campo Grande\",\n  \"localidade\": \"Rio de Janeiro\",\n  \"uf\": \"RJ\",\n  \"ibge\": \"3304557\",\n  \"gia\": \"\",\n  \"ddd\": \"21\",\n  \"siafi\": \"6001\"\n}"

	var viaCEPChannel = make(chan string)
	viaCep := NewViaCep(viaCEPChannel)

	cep, err := viaCep.GetCEP("23017-335")

	if err != nil {
		t.Errorf("Error on get cep")
	}

	if cep != cep_expected {
		t.Errorf("Cep is not valid")
	}

}
