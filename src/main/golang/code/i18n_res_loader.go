package code

import (
	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
)

type i18nStringsLoader struct {
	ac application.Context
}

func (inst *i18nStringsLoader) load(path string) map[string]string {
	dst := make(map[string]string)
	mods := inst.ac.GetModules()
	for _, mod := range mods {
		inst.loadFromModule(mod, path, dst)
	}
	return dst
}

func (inst *i18nStringsLoader) loadFromModule(mod application.Module, path string, dst map[string]string) {

	if mod == nil {
		return
	}

	res := mod.Resources()
	if res == nil {
		return
	}

	text, err := res.ReadText(path)
	if err != nil {
		return
	}

	props, err := properties.Parse(text, nil)
	if err != nil {
		return
	}

	props.Export(dst)
}
