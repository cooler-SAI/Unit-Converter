package converters

import (
	"fmt"
	"strings"
)

// lengthInMeters holds conversion factors from each unit to meters.
var lengthInMeters = map[string]float64{
	"millimeter": 0.001,
	"centimeter": 0.01,
	"meter":      1.0,
	"kilometer":  1000.0,
	"inch":       0.0254,
	"foot":       0.3048,
	"yard":       0.9144,
	"mile":       1609.344,
}

// ConvertLength converts a value from one length unit to another.
func ConvertLength(value float64, fromUnit, toUnit string) (float64, error) {
	fromUnit = strings.ToLower(fromUnit)
	toUnit = strings.ToLower(toUnit)

	fromFactor, okFrom := lengthInMeters[fromUnit]
	toFactor, okTo := lengthInMeters[toUnit]
	if !okFrom || !okTo {
		return 0, fmt.Errorf("unsupported length unit")
	}

	// Convert value to meters first
	valueInMeters := value * fromFactor
	return valueInMeters / toFactor, nil
}
