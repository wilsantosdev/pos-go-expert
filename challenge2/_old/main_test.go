package main

import "testing"

func TestValidateCepFormat(t *testing.T) {
	validCEP := validateCepFormat("23017-335")

	if validCEP != true {
		t.Errorf("CEP is not valid")
	}
}
