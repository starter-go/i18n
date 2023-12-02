package code

import (
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
		Test:    inst.test,
	})
	return list
}

func (inst *TestI18nRes) test() error {
	str, err := inst.Service.Default().GetString("a.b.c")
	if err != nil {
		return err
	}
	vlog.Debug(str)
	return nil
}
