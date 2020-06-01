package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/helpers/forms"
	"github.com/gobuffalo/packr/v2"
	"github.com/henry-jackson/challenges/sales_admin/sales_app/models"
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
			"orderTotal": func(price float64, qty int) string {
				total := float64(qty) * price
				return fmt.Sprintf("$%6.2f", total)
			},
			"totalRevenue": func(orders []models.Order) string {
				sum := 0.00
				for _, order := range orders {
					sum += float64(order.Quantity) * order.Product.Price
				}
				return fmt.Sprintf("$%6.2f", sum)
			},
			"formatTime": func(t time.Time) string {
				return t.Format("January 2, 2006")
			},
			"formatMoney": func(amt float64) string {
				return fmt.Sprintf("$%6.2f", amt)
			},
		},
	})
}
