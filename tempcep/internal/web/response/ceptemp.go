package response

type cepTempResponse struct {
	Celsius   float64 `json:"temp_C"`
	Farenheit float64 `json:"temp_F"`
	Kelvin    float64 `json:"temp_K"`
}

func NewCepTempResponse(celsius, farenheit, kelvin float64) cepTempResponse {
	return cepTempResponse{
		Celsius:   celsius,
		Farenheit: farenheit,
		Kelvin:    kelvin,
	}
}
