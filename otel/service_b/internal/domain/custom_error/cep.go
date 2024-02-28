package customerror

type CEPInvalidFormat struct{}

func (e CEPInvalidFormat) Error() string {
	return "CEP invalid format"
}
