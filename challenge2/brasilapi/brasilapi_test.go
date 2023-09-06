package brasilapi

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
		a    brasilApi
		args args
		want string
	}{
		{
			name: "TestMakeUrl",
			a: brasilApi{
				channel: make(chan string),
			},
			args: args{
				cep: "01001-000",
			},
			want: "https://brasilapi.com.br/api/cep/v1/01001-000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.makeUrl(tt.args.cep); got != tt.want {
				t.Errorf("brasilApi.makeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCEP(t *testing.T) {

	cep_expected := `"cep":"01001000"`
	state_expected := `"state":"SP"`
	city_expected := `"city":"São Paulo"`
	neighborhood_expected := `"neighborhood":"Sé"`
	street_expected := `"street":"Praça da Sé`

	var brasilAPIChannel = make(chan string)
	brasilAPISerice := NewBrasilApi(&brasilAPIChannel)

	go brasilAPISerice.GetCEP("01001-000")

	cep := <-brasilAPIChannel

	if !strings.Contains(cep, cep_expected) {
		t.Errorf("brasilApi.GetCEP() = %v, want %v", cep, cep_expected)
	}

	if !strings.Contains(cep, state_expected) {
		t.Errorf("brasilApi.GetCEP() = %v, want %v", cep, state_expected)
	}

	if !strings.Contains(cep, city_expected) {
		t.Errorf("brasilApi.GetCEP() = %v, want %v", cep, city_expected)
	}

	if !strings.Contains(cep, neighborhood_expected) {
		t.Errorf("brasilApi.GetCEP() = %v, want %v", cep, neighborhood_expected)
	}

	if !strings.Contains(cep, street_expected) {
		t.Errorf("brasilApi.GetCEP() = %v, want %v", cep, street_expected)
	}

}
