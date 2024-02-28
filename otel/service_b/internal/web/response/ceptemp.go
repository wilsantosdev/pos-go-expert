package response

type cepTempResponse struct {
	City      string  `json:"city"`
	Celsius   float64 `json:"temp_C"`
	Farenheit float64 `json:"temp_F"`
	Kelvin    float64 `json:"temp_K"`
}

func NewCepTempResponse(city string, celsius, farenheit, kelvin float64) cepTempResponse {
	return cepTempResponse{
		City:      city,
		Celsius:   celsius,
		Farenheit: farenheit,
		Kelvin:    kelvin,
	}
}
