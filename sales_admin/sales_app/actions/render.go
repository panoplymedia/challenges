package actions

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/helpers/forms"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// forms.FormForKey:  forms.FormFor,
			forms.FormKey: forms.Form,
			"activeClass": func(n string, help plush.HelperContext) string {
				if p, ok := help.Value("current_route").(buffalo.RouteInfo); ok {
					if strings.Contains(p.PathName, n) {
						return "active"
					}
				}
				return ""
			},
		},
	})
}
