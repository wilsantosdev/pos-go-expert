package customerror

type CEPNotFound struct{}

func (e CEPNotFound) Error() string {
	return "can not find zipcode"
}
