package customerror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCep(t *testing.T) {
	t.Parallel()
	t.Run("CEPInvalidFormat", func(t *testing.T) {
		t.Parallel()
		cep := CEPInvalidFormat{}
		assert.Equal(t, "CEP invalid format", cep.Error())
	})
}
