package code

import (
	"sort"

	"github.com/starter-go/i18n"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// TestI18nRes ...
type TestI18nRes struct {

	//starter:component

	Service i18n.Service //starter:inject("#")

}

func (inst *TestI18nRes) _impl() units.Units { return inst }

// Units ...
func (inst *TestI18nRes) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:    "test-i18n-res",
		Enabled: true,
		Test:    inst.testGetString,
	})

	list = append(list, &units.Registration{
		Name:    "test-i18n-list-strings",
		Enabled: true,
		Test:    inst.testListStrings,
	})

	return list
}

func (inst *TestI18nRes) testGetString() error {

	keys := make([]string, 0)
	keys = append(keys, "a.b.c")
	keys = append(keys, "i.j.k")
	keys = append(keys, "x.y.z")

	for i, key := range keys {
		val := inst.Service.Default().String(key)
		vlog.Info("test[%d] String(%s) = %s", i, key, val)
	}
	return nil
}

func (inst *TestI18nRes) testListStrings() error {

	// res := inst.Service.Default()
	res := inst.Service.GetResources("default", "en_US", "zh_CN")

	names := res.Names()
	sort.Strings(names)

	for i, name := range names {
		val := res.String(name)
		vlog.Info("list strings[%d]: string(%s) = %s", i, name, val)
	}

	return nil
}
