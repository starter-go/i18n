package main

import (
	"os"

	"github.com/starter-go/i18n/modules/i18n"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := i18n.ModuleForTest()
	c := new(units.Context)

	c.Arguments = a
	c.Module = m
	c.UsePanic = true

	units.Run(c)
}
