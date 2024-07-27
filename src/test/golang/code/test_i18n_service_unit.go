package code

import (
	"github.com/starter-go/i18n"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestI18nServiceUnit ...
type TestI18nServiceUnit struct {

	//starter:component

	Service i18n.Service //starter:inject("#")

}

func (inst *TestI18nServiceUnit) _impl() units.Units {
	return inst
}

// Units ...
func (inst *TestI18nServiceUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:    CaseTestServiceAvailable,
		Test:    inst.testAvailable,
		Enabled: true,
	})

	list = append(list, &units.Registration{
		Name:    CaseTestServiceDefault,
		Test:    inst.testDefault,
		Enabled: true,
	})

	list = append(list, &units.Registration{
		Name:    CaseTestServiceGetResources,
		Test:    inst.testGetResources,
		Enabled: true,
	})

	return list
}

func (inst *TestI18nServiceUnit) testAvailable() error {
	all := inst.Service.Available()
	for i, lang := range all {
		vlog.Info("i18n.langs[%d] = %s", i, lang)
	}
	return nil
}

func (inst *TestI18nServiceUnit) testDefault() error {
	def := inst.Service.Default()
	list := def.Names()
	for i, name := range list {
		vlog.Info("i18n.names[%d].name = %s", i, name)
	}
	return nil
}

func (inst *TestI18nServiceUnit) testGetResources() error {
	res := inst.Service.GetResources("zh_cn", "default", "fr_fr", "en_us")
	key := "strings.a"
	val, _ := res.GetString(key)
	vlog.Info("TestI18nServiceUnit.testGetResources: %s = %s", key, val)
	return nil
}
