package converters

import (
	"fmt"
	"strings"
)

// ConvertTemperature converts a temperature value from one scale to another (Celsius, Fahrenheit, Kelvin).
func ConvertTemperature(value float64, fromUnit, toUnit string) (float64, error) {
	fromUnit = strings.ToLower(fromUnit)
	toUnit = strings.ToLower(toUnit)

	// Convert from fromUnit to Celsius
	var valueInCelsius float64

	switch fromUnit {
	case "celsius":
		valueInCelsius = value
	case "fahrenheit":
		// (F - 32) * 5/9
		valueInCelsius = (value - 32) * 5.0 / 9.0
	case "kelvin":
		// K - 273.15
		valueInCelsius = value - 273.15
	default:
		return 0, fmt.Errorf("unsupported temperature unit")
	}

	// Convert from Celsius to toUnit
	switch toUnit {
	case "celsius":
		return valueInCelsius, nil
	case "fahrenheit":
		// (C * 9/5) + 32
		return (valueInCelsius * 9.0 / 5.0) + 32, nil
	case "kelvin":
		// C + 273.15
		return valueInCelsius + 273.15, nil
	default:
		return 0, fmt.Errorf("unsupported temperature unit")
	}
}
