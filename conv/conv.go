package conv

import "math"

func FloatRound(value float64) float64 { //Dette er funksjon for å runde av
	return math.Round(value*100) / 100
}

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	celsius := (value - 32) * (5.0 / 9)
	return FloatRound(celsius)
}

// Konverterer Celsius til Farhenheit
func CelsiusToFahrenheit(value float64) float64 { //Farhrenheit = Celsius*(9/5) + 32
	farhenheit := (value)*(9/5.0) + 32
	return FloatRound(farhenheit)
}

//Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 { //Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
	farhenheit := (value-273.15)*(9/5.0) + 32
	return FloatRound(farhenheit)
}

//Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 { //(Farhenheit - 32) * (5/9) + 273.15
	kelvin := (value-32)*(5.0/9) + 273.15
	return FloatRound(kelvin)
}

//Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 { //Celsius = Kelvin - 273.15
	celsius := value - 273.15
	return FloatRound(celsius)
}

//Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 { //Kelvin = Celsius + 273.15
	kelvin := value + 273.15
	return FloatRound(kelvin)
}
