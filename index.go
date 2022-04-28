package view_default

import (
	"github.com/chefsgo/view"
)

func Driver() view.Driver {
	return &defaultDriver{}
}

func init() {
	view.Register("default", Driver())
}
