package main

import (
	"testing"

	"github.com/starter-go/application"
	"github.com/starter-go/i18n/modules/i18n"
	"github.com/starter-go/i18n/src/test/golang/code"
	"github.com/starter-go/units"
)

func getTestModule() application.Module {
	return i18n.ModuleForTest()
}

func TestI18nServiceAvailable(t *testing.T) {
	args := []string{
		"--debug.enabled=1",
	}
	mod := getTestModule()
	cfg := &units.Config{
		Args:   args,
		Cases:  code.CaseTestServiceAvailable,
		Module: mod,
		T:      t,
	}
	units.Run(cfg)
}

func TestI18nServiceDefault(t *testing.T) {
	args := []string{
		"--debug.enabled=1",
	}
	mod := getTestModule()
	cfg := &units.Config{
		Args:   args,
		Cases:  code.CaseTestServiceDefault,
		Module: mod,
		T:      t,
	}
	units.Run(cfg)
}

func TestI18nServiceGetResources(t *testing.T) {
	args := []string{
		"--debug.enabled=1",
	}
	mod := getTestModule()
	cfg := &units.Config{
		Args:   args,
		Cases:  code.CaseTestServiceGetResources,
		Module: mod,
		T:      t,
	}
	units.Run(cfg)
}
