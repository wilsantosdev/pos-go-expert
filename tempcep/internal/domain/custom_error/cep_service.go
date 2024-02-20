package customerror

type CEPNotFound struct{}

func (e CEPNotFound) Error() string {
	return "CEP not found"
}
