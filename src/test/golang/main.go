package main

import (
	"os"

	"github.com/starter-go/i18n/modules/i18n"
	"github.com/starter-go/units"
)

func main() {
	m := i18n.ModuleForTest()
	units.NewRunner().Dependencies(m).Run(os.Args)
}
