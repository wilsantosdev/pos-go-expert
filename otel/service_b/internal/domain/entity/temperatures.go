package entity

type Temperatures struct {
	city      string
	celsius   float64
	farenheit float64
	kelvin    float64
}

func NewTemperatures(city string, celsius float64, farenheit float64, kelvin float64) Temperatures {
	return Temperatures{
		city:      city,
		celsius:   celsius,
		farenheit: farenheit,
		kelvin:    kelvin,
	}
}

func (t Temperatures) Celsius() float64 {
	return t.celsius
}

func (t Temperatures) Farenheit() float64 {
	return t.farenheit
}

func (t Temperatures) Kelvin() float64 {
	return t.kelvin
}

func (t Temperatures) City() string {
	return t.city
}
