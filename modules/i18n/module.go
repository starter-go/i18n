package i18n

import (
	"github.com/starter-go/application"
	"github.com/starter-go/i18n"
	"github.com/starter-go/i18n/gen/main4i18n"
	"github.com/starter-go/i18n/gen/test4i18n"
	"github.com/starter-go/units/modules/units"
)

// Module 导出模块
func Module() application.Module {
	mb := i18n.NewMainModule()
	mb.Components(main4i18n.ExportComponents)
	return mb.Create()
}

// ModuleForTest 导出模块
func ModuleForTest() application.Module {

	parent := Module()

	mb := i18n.NewTestModule()
	mb.Components(test4i18n.ExportComponents)

	mb.Depend(parent)
	mb.Depend(units.Module())

	return mb.Create()
}
