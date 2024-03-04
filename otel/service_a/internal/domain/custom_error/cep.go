package customerror

type CEPInvalidFormat struct{}

func (e CEPInvalidFormat) Error() string {
	return "invalid zipcode"
}

type CEPNotFound struct{}

func (e CEPNotFound) Error() string {
	return "can not find zipcode"
}
