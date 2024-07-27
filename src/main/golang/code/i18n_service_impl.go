package code

import (
	"github.com/starter-go/base/util"
	"github.com/starter-go/i18n"
)

// I18nServiceImpl ...
type I18nServiceImpl struct {

	//starter:component
	_as func(i18n.Service) //starter:as("#")

	Providers []i18n.LanguageProvider //starter:inject(".")

	cache *cache
}

func (inst *I18nServiceImpl) _impl() i18n.Service {
	return inst
}

// GetResources ...
func (inst *I18nServiceImpl) GetResources(l ...i18n.Language) i18n.Resources {
	c := inst.getCache()
	return c.makeResourceChain(l...)
}

// Default ...
func (inst *I18nServiceImpl) Default() i18n.Resources {
	c := inst.getCache()
	return c.defaultRes
}

// Available ...
func (inst *I18nServiceImpl) Available() []i18n.Language {
	c := inst.getCache()
	return c.available
}

func (inst *I18nServiceImpl) getCache() *cache {
	c := inst.cache
	if c != nil {
		return c
	}
	c = inst.loadCache()
	inst.cache = c
	return c
}

func (inst *I18nServiceImpl) loadCache() *cache {
	c := &cache{service: inst}
	c.load()
	return c
}

////////////////////////////////////////////////////////////////////////////////

type cache struct {
	service *I18nServiceImpl

	available  []i18n.Language
	all        []*i18n.LanguageRegistration
	defaultRes i18n.Resources
}

func (inst *cache) load() {
	inst.loadAll()
	inst.loadAvailable()
	inst.loadDefault()
}

func (inst *cache) loadAll() {

	src := inst.service.Providers
	tmp := make([]*i18n.LanguageRegistration, 0)
	dst := make([]*i18n.LanguageRegistration, 0)

	for _, r1 := range src {
		list := r1.Languages()
		tmp = append(tmp, list...)
	}
	for _, item := range tmp {
		if item == nil {
			continue
		}
		if !item.Enabled {
			continue
		}
		dst = append(dst, item)
	}

	s := util.Sorter{
		OnLen:  func() int { return len(dst) },
		OnLess: func(i1, i2 int) bool { return dst[i1].Priority > dst[i2].Priority },
		OnSwap: func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] },
	}
	s.Sort()

	inst.all = dst
}

func (inst *cache) loadAvailable() {
	src := inst.all
	dst := make([]i18n.Language, 0)
	for _, item := range src {
		dst = append(dst, item.Language)
	}
	inst.available = dst
}

func (inst *cache) loadDefault() {
	chain := inst.makeResourceChain(inst.available...)
	inst.defaultRes = chain
}

func (inst *cache) makeResourceChain(list ...i18n.Language) i18n.Resources {
	src := inst.all
	table := make(map[i18n.Language]*i18n.LanguageRegistration)
	for _, item := range src {
		key := item.Language.Normalize()
		table[key] = item
	}
	b := &i18nResourceChainBuilder{}
	for _, name := range list {
		key := name.Normalize()
		item := table[key]
		if item == nil {
			continue
		}
		b.add(item)
	}
	return b.create()
}
