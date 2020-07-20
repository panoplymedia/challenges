package service

import (
	html_template "html/template"
	"io"

	"github.com/labstack/echo"
)

// Template is an implementation of the echo.Renderer so that we can html.Template.
type template html_template.Template

// Render renders the html templates to display on the browser.
func (t *template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return (*html_template.Template)(t).ExecuteTemplate(w, name, data)
}
