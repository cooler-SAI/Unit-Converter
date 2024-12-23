package converters

import (
	"fmt"
	"strings"
)

// weightInGrams holds conversion factors from each unit to grams.
var weightInGrams = map[string]float64{
	"milligram": 0.001,
	"gram":      1.0,
	"kilogram":  1000.0,
	"ounce":     28.349523125,
	"pound":     453.59237,
}

// ConvertWeight converts a value from one weight unit to another.
func ConvertWeight(value float64, fromUnit, toUnit string) (float64, error) {
	fromUnit = strings.ToLower(fromUnit)
	toUnit = strings.ToLower(toUnit)

	fromFactor, okFrom := weightInGrams[fromUnit]
	toFactor, okTo := weightInGrams[toUnit]
	if !okFrom || !okTo {
		return 0, fmt.Errorf("unsupported weight unit")
	}

	// Convert value to grams first
	valueInGrams := value * fromFactor
	// Then convert from grams to the target unit
	return valueInGrams / toFactor, nil
}
