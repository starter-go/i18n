package code

import (
	"fmt"

	"github.com/starter-go/i18n"
)

type i18nResourceChainBuilder struct {
	items []*i18n.LanguageRegistration
}

func (inst *i18nResourceChainBuilder) add(lr *i18n.LanguageRegistration) {
	inst.items = append(inst.items, lr)
}

func (inst *i18nResourceChainBuilder) create() i18n.Resources {
	items := inst.items
	var p i18n.Resources
	p = &ending{}
	for _, item := range items {
		n := &node{}
		n.next = p
		n.target = item.Resources
		n.info = *item
		p = n
	}
	return p
}

////////////////////////////////////////////////////////////////////////////////

type node struct {
	next   i18n.Resources
	info   i18n.LanguageRegistration
	target i18n.Resources
}

func (inst *node) _impl() i18n.Resources { return inst }

func (inst *node) ReadText(path string) (string, error) {
	text, err := inst.target.ReadText(path)
	if err == nil {
		return text, nil
	}
	return inst.next.ReadText(path)
}

func (inst *node) ReadBinary(path string) ([]byte, error) {
	data, err := inst.target.ReadBinary(path)
	if err == nil {
		return data, nil
	}
	return inst.next.ReadBinary(path)
}

func (inst *node) ReadJSON(path string, root any) error {
	err := inst.target.ReadJSON(path, root)
	if err == nil {
		return nil
	}
	return inst.next.ReadJSON(path, root)
}

func (inst *node) GetString(name string) (string, error) {
	str, err := inst.target.GetString(name)
	if err == nil {
		return str, nil
	}
	return inst.next.GetString(name)
}

////////////////////////////////////////////////////////////////////////////////

type ending struct{}

func (inst *ending) _impl() i18n.Resources { return inst }

func (inst *ending) ReadText(path string) (string, error) {
	err := inst.read(path)
	return "", err
}

func (inst *ending) ReadBinary(path string) ([]byte, error) {
	err := inst.read(path)
	return nil, err
}

func (inst *ending) ReadJSON(path string, root any) error {
	return inst.read(path)
}

func (inst *ending) read(path string) error {
	return fmt.Errorf("no i18n resource at path [%s]", path)
}

func (inst *ending) GetString(name string) (string, error) {
	err := fmt.Errorf("no i18n string with name [%s]", name)
	return "", err
}

////////////////////////////////////////////////////////////////////////////////
