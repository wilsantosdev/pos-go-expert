package customerror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCepService(t *testing.T) {
	t.Parallel()
	t.Run("CEPNotFound", func(t *testing.T) {
		t.Parallel()
		cep := CEPNotFound{}
		assert.Equal(t, "CEP not found", cep.Error())
	})
}
