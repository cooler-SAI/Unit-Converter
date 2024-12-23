package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"Unit-Converter/converters"
)

// LengthHandler displays the form and processes length conversion.
func LengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Retrieve form values
		valueStr := r.FormValue("value")
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")

		// Convert string to float
		valueFloat, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			ShowError(w, fmt.Sprintf("Invalid number input: %v", err))
			return
		}

		// Call the converter function
		result, err := converters.ConvertLength(valueFloat, fromUnit, toUnit)
		if err != nil {
			ShowError(w, err.Error())
			return
		}

		// Build result form
		resultHTML := fmt.Sprintf(`
            <form method="POST" action="/length">
                <label>Value:</label><br>
                <input type="text" name="value" value="%s"><br>

                <label>From Unit (e.g. meter):</label><br>
                <input type="text" name="fromUnit" value="%s"><br>

                <label>To Unit (e.g. foot):</label><br>
                <input type="text" name="toUnit" value="%s"><br>

                <input type="submit" value="Convert">
            </form>

            <div class="result">Result: %.4f %s</div>
        `, valueStr, fromUnit, toUnit, result, toUnit)

		data := PageData{Content: template.HTML(resultHTML)}
		_ = tmpl.Execute(w, data)
	} else {
		// Display a blank form
		formHTML := `
            <form method="POST" action="/length">
                <label>Value:</label><br>
                <input type="text" name="value" placeholder="Enter value"><br>

                <label>From Unit (e.g. meter):</label><br>
                <input type="text" name="fromUnit" placeholder="meter"><br>

                <label>To Unit (e.g. foot):</label><br>
                <input type="text" name="toUnit" placeholder="foot"><br>

                <input type="submit" value="Convert">
            </form>
        `
		data := PageData{Content: template.HTML(formHTML)}
		_ = tmpl.Execute(w, data)
	}
}
