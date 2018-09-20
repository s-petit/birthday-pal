package main

import (
	"github.com/s-petit/birthday-pal/app"
	"github.com/s-petit/birthday-pal/cli"
	"github.com/s-petit/birthday-pal/system"
)

func main() {
	cli.Mowcli(app.BirthdayPal{}, system.RealSystem{})
}
