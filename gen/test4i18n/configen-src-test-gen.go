package test4i18n
import (
    p52d293b71 "github.com/starter-go/i18n"
    p856ff22f9 "github.com/starter-go/i18n/src/test/golang/code"
     "github.com/starter-go/application"
)

// type p856ff22f9.TestI18nRes in package:github.com/starter-go/i18n/src/test/golang/code
//
// id:com-856ff22f92454c1a-code-TestI18nRes
// class:
// alias:
// scope:singleton
//
type p856ff22f92_code_TestI18nRes struct {
}

func (inst* p856ff22f92_code_TestI18nRes) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-856ff22f92454c1a-code-TestI18nRes"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p856ff22f92_code_TestI18nRes) new() any {
    return &p856ff22f9.TestI18nRes{}
}

func (inst* p856ff22f92_code_TestI18nRes) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p856ff22f9.TestI18nRes)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p856ff22f92_code_TestI18nRes) getService(ie application.InjectionExt)p52d293b71.Service{
    return ie.GetComponent("#alias-52d293b7197ae4adfab092ade2e718a2-Service").(p52d293b71.Service)
}



// type p856ff22f9.TestI18nServiceUnit in package:github.com/starter-go/i18n/src/test/golang/code
//
// id:com-856ff22f92454c1a-code-TestI18nServiceUnit
// class:
// alias:
// scope:singleton
//
type p856ff22f92_code_TestI18nServiceUnit struct {
}

func (inst* p856ff22f92_code_TestI18nServiceUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-856ff22f92454c1a-code-TestI18nServiceUnit"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p856ff22f92_code_TestI18nServiceUnit) new() any {
    return &p856ff22f9.TestI18nServiceUnit{}
}

func (inst* p856ff22f92_code_TestI18nServiceUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p856ff22f9.TestI18nServiceUnit)
	nop(ie, com)

	
    com.Service = inst.getService(ie)


    return nil
}


func (inst*p856ff22f92_code_TestI18nServiceUnit) getService(ie application.InjectionExt)p52d293b71.Service{
    return ie.GetComponent("#alias-52d293b7197ae4adfab092ade2e718a2-Service").(p52d293b71.Service)
}


