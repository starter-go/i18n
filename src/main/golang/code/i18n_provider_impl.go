package code

import (
	"encoding/json"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/base/safe"
	"github.com/starter-go/i18n"
	"github.com/starter-go/vlog"
)

// ProviderImpl ...
type ProviderImpl struct {

	//starter:component
	_as func(i18n.LanguageProvider) //starter:as(".")

	AC       application.Context //starter:inject("context")
	LangList string              //starter:inject("${i18n.available.languages}")

}

func (inst *ProviderImpl) _impl() i18n.LanguageProvider {
	return inst
}

// Languages ...
func (inst *ProviderImpl) Languages() []*i18n.LanguageRegistration {
	namelist := inst.listNames()
	dst := make([]*i18n.LanguageRegistration, 0)
	for i, name := range namelist {
		item := inst.loadLanguage(i, name)
		dst = append(dst, item)
	}
	return dst
}

func (inst *ProviderImpl) listNames() []i18n.Language {
	str := inst.LangList
	src := strings.Split(str, ",")
	dst := make([]i18n.Language, 0)
	for _, item := range src {
		name := i18n.Language(item)
		name = name.Normalize()
		if name != "" {
			dst = append(dst, name)
		}
	}
	return dst
}

func (inst *ProviderImpl) loadLanguage(index int, l i18n.Language) *i18n.LanguageRegistration {
	lr := &i18n.LanguageRegistration{}
	lr.Language = l.Normalize()
	lr.Enabled = true
	lr.Priority = 1000 - index
	lr.Resources = inst.makeRes(l)
	return lr
}

func (inst *ProviderImpl) makeRes(l i18n.Language) i18n.Resources {
	src := inst.AC.GetResources()
	res := &resourcesImpl{
		lang: l,
		src:  src,
		ac:   inst.AC,
	}
	return res
}

////////////////////////////////////////////////////////////////////////////////

type resourcesImpl struct {
	ac      application.Context
	src     resources.Table
	lang    i18n.Language
	strings properties.Table
}

func (inst *resourcesImpl) _impl() i18n.Resources { return inst }

func (inst *resourcesImpl) computePath(path string) string {
	lang := inst.lang.String()
	return "res:///i18n/" + lang + "/" + path
}

func (inst *resourcesImpl) Names() []string {
	return inst.getStrings().Names()
}

func (inst *resourcesImpl) ReadText(path string) (string, error) {
	p2 := inst.computePath(path)
	return inst.src.ReadText(p2)
}

func (inst *resourcesImpl) ReadBinary(path1 string) ([]byte, error) {
	path2 := inst.computePath(path1)
	bin, err := inst.src.ReadBinary(path2)
	if err != nil {
		vlog.Warn(err.Error())
	}
	return bin, err
}

func (inst *resourcesImpl) ReadJSON(path string, root any) error {
	data, err := inst.ReadBinary(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, root)
}

func (inst *resourcesImpl) GetString(name string) (string, error) {
	table := inst.getStrings()
	return table.GetPropertyRequired(name)
}

func (inst *resourcesImpl) String(name string) string {
	str, err := inst.GetString(name)
	if err != nil {
		str = "[" + err.Error() + "]"
	}
	return str
}

func (inst *resourcesImpl) getStrings() properties.Table {
	table := inst.strings
	if table != nil {
		return table
	}
	table = inst.loadStrings()
	inst.strings = table
	return table
}

func (inst *resourcesImpl) loadStrings() properties.Table {

	path1 := "strings.properties"
	path2 := inst.computePath(path1)

	loader := new(i18nStringsLoader)
	loader.ac = inst.ac
	src := loader.load(path2)

	mode := safe.Safe()
	dst := properties.NewTable(mode)
	dst.Import(src)
	return dst
}
