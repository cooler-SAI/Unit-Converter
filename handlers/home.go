package handlers

import (
	"net/http"
)

// HomeHandler serves the main page.
func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	data := PageData{
		Content: "<p>Welcome to the Unit Converter! Choose a conversion type above.</p>",
	}
	_ = tmpl.Execute(w, data)
}
