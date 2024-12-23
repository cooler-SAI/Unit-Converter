package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"Unit-Converter/converters"
)

// TemperatureHandler displays the form and processes temperature conversion.
func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")

		valueFloat, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			ShowError(w, fmt.Sprintf("Invalid number input: %v", err))
			return
		}

		result, err := converters.ConvertTemperature(valueFloat, fromUnit, toUnit)
		if err != nil {
			ShowError(w, err.Error())
			return
		}

		resultHTML := fmt.Sprintf(`
            <form method="POST" action="/temperature">
                <label>Value:</label><br>
                <input type="text" name="value" value="%s"><br>

                <label>From Unit (Celsius/Fahrenheit/Kelvin):</label><br>
                <input type="text" name="fromUnit" value="%s"><br>

                <label>To Unit (Celsius/Fahrenheit/Kelvin):</label><br>
                <input type="text" name="toUnit" value="%s"><br>

                <input type="submit" value="Convert">
            </form>

            <div class="result">Result: %.2f %s</div>
        `, valueStr, fromUnit, toUnit, result, toUnit)

		data := PageData{Content: template.HTML(resultHTML)}
		_ = tmpl.Execute(w, data)
	} else {
		formHTML := `
            <form method="POST" action="/temperature">
                <label>Value:</label><br>
                <input type="text" name="value" placeholder="Enter value"><br>

                <label>From Unit (Celsius/Fahrenheit/Kelvin):</label><br>
                <input type="text" name="fromUnit" placeholder="Celsius"><br>

                <label>To Unit (Celsius/Fahrenheit/Kelvin):</label><br>
                <input type="text" name="toUnit" placeholder="Fahrenheit"><br>

                <input type="submit" value="Convert">
            </form>
        `
		data := PageData{Content: template.HTML(formHTML)}
		_ = tmpl.Execute(w, data)
	}
}
