package viacep

import (
	"strings"
	"testing"
)

func TestMakeUrl(t *testing.T) {
	type args struct {
		cep string
	}
	tests := []struct {
		name string
		a    viaCep
		args args
		want string
	}{
		{
			name: "TestMakeUrl",
			a: viaCep{
				channel: make(chan string),
			},
			args: args{
				cep: "01001000",
			},
			want: "https://viacep.com.br/ws/01001000/json/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.makeUrl(tt.args.cep); got != tt.want {
				t.Errorf("viaCep.makeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetCEP(t *testing.T) {

	cep_expected := `"cep": "01001-000"`
	logradouro_expected := `"logradouro": "Praça da Sé"`
	complemento_expected := `"complemento": "lado ímpar"`
	bairro_expected := `"bairro": "Sé"`
	localidade_expected := `"localidade": "São Paulo"`
	uf_expected := `"uf": "SP"`
	ibge_expected := `"ibge": "3550308"`
	gia_expected := `"gia": "1004"`
	ddd_expected := `"ddd": "11"`
	siafi_expected := `"siafi": "7107"`

	var viaCEPChannel = make(chan string)
	viaCepService := NewViaCep(&viaCEPChannel)

	go viaCepService.GetCEP("01001-000")

	cep := <-viaCEPChannel

	if !strings.Contains(cep, cep_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, cep_expected)
	}

	if !strings.Contains(cep, logradouro_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, logradouro_expected)
	}

	if !strings.Contains(cep, complemento_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, complemento_expected)
	}

	if !strings.Contains(cep, bairro_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, bairro_expected)
	}

	if !strings.Contains(cep, localidade_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, localidade_expected)
	}

	if !strings.Contains(cep, uf_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, `"uf": "SP"`)
	}

	if !strings.Contains(cep, ibge_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, ibge_expected)
	}

	if !strings.Contains(cep, gia_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, gia_expected)
	}

	if !strings.Contains(cep, ddd_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, ddd_expected)
	}

	if !strings.Contains(cep, siafi_expected) {
		t.Errorf("viaCep.GetCEP() = %v, want %v", cep, siafi_expected)
	}

}
