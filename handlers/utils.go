package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// PageData holds any dynamic content to be inserted into the template.
type PageData struct {
	Content template.HTML
}

// tmpl is our main parsed template. We parse "templates/index.html" once on init.
var tmpl *template.Template

// init runs automatically when the package is loaded. We parse the template here.
func init() {
	// Construct the full path to the template file (assuming we run from project root)
	templatePath := filepath.Join("templates", "index.html")

	tmpl = template.Must(template.ParseFiles(templatePath))
}

func ShowError(w http.ResponseWriter, errMsg string) {
	data := PageData{
		Content: template.HTML(fmt.Sprintf("<p class='error'>Error: %s</p>", errMsg)),
	}
	_ = tmpl.Execute(w, data)
}
