package main

import (
	"github.com/s-petit/birthday-pal/bpal"
	"github.com/s-petit/birthday-pal/cli"
	"github.com/s-petit/birthday-pal/system"
)

func main() {
	cli.Mowcli(bpal.BirthdayPal{}, system.RealSystem{})
}
