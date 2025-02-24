package main

import (
	"fmt"
)

type Celsium = float64
type Kelvin = float64

func CelsiumToKelvin(cel Celsium) Kelvin {
	return Kelvin(cel + 273.15)
}

func KelvinToCelsium(kel Kelvin) Celsium {
	return Celsium(kel - 273.15)
}

func SwapTemperatures(kelPtr *Kelvin, celPtr *Celsium) {
	temp_kel := KelvinToCelsium(*kelPtr)
	*kelPtr = CelsiumToKelvin(*celPtr)
	*celPtr = temp_kel
}

func main() {
	kelPtr := new(Kelvin)
	*kelPtr = 315.92
	celPtr := new(Celsium)
	*celPtr = 20.0

	SwapTemperatures(kelPtr, celPtr)
	fmt.Printf("%.2f\n", *celPtr)
	fmt.Printf("%.2f\n", *kelPtr)
}
