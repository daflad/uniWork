package models

//CtoF convert Cel to Far
//Celsius (°C) times 9/5 plus 32:
func CtoF(input float64) float64 {
	return (input * 9 / 5) + 32
}

//FtoC convert Far to Cel
//Fahrenheit (°F) minus 32, times 5/9
func FtoC(input float64) float64 {
	return (input - 32) * 5 / 9
}
