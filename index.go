package view_default

import (
	"github.com/chefsgo/chef"
)

func Driver() chef.ViewDriver {
	return &defaultViewDriver{}
}

func init() {
	chef.Register("default", Driver())
}
