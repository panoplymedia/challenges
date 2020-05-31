package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/henry-jackson/challenges/sales_admin/sales_app/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
