package main4i18n
import (
    p0ef6f2938 "github.com/starter-go/application"
    p52d293b71 "github.com/starter-go/i18n"
    p2afe58e9d "github.com/starter-go/i18n/src/main/golang/code"
     "github.com/starter-go/application"
)

// type p2afe58e9d.ProviderImpl in package:github.com/starter-go/i18n/src/main/golang/code
//
// id:com-2afe58e9d62bf3f3-code-ProviderImpl
// class:class-52d293b7197ae4adfab092ade2e718a2-LanguageProvider
// alias:
// scope:singleton
//
type p2afe58e9d6_code_ProviderImpl struct {
}

func (inst* p2afe58e9d6_code_ProviderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2afe58e9d62bf3f3-code-ProviderImpl"
	r.Classes = "class-52d293b7197ae4adfab092ade2e718a2-LanguageProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2afe58e9d6_code_ProviderImpl) new() any {
    return &p2afe58e9d.ProviderImpl{}
}

func (inst* p2afe58e9d6_code_ProviderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2afe58e9d.ProviderImpl)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.LangList = inst.getLangList(ie)


    return nil
}


func (inst*p2afe58e9d6_code_ProviderImpl) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p2afe58e9d6_code_ProviderImpl) getLangList(ie application.InjectionExt)string{
    return ie.GetString("${i18n.available.languages}")
}



// type p2afe58e9d.I18nServiceImpl in package:github.com/starter-go/i18n/src/main/golang/code
//
// id:com-2afe58e9d62bf3f3-code-I18nServiceImpl
// class:
// alias:alias-52d293b7197ae4adfab092ade2e718a2-Service
// scope:singleton
//
type p2afe58e9d6_code_I18nServiceImpl struct {
}

func (inst* p2afe58e9d6_code_I18nServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2afe58e9d62bf3f3-code-I18nServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-52d293b7197ae4adfab092ade2e718a2-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2afe58e9d6_code_I18nServiceImpl) new() any {
    return &p2afe58e9d.I18nServiceImpl{}
}

func (inst* p2afe58e9d6_code_I18nServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2afe58e9d.I18nServiceImpl)
	nop(ie, com)

	
    com.Providers = inst.getProviders(ie)


    return nil
}


func (inst*p2afe58e9d6_code_I18nServiceImpl) getProviders(ie application.InjectionExt)[]p52d293b71.LanguageProvider{
    dst := make([]p52d293b71.LanguageProvider, 0)
    src := ie.ListComponents(".class-52d293b7197ae4adfab092ade2e718a2-LanguageProvider")
    for _, item1 := range src {
        item2 := item1.(p52d293b71.LanguageProvider)
        dst = append(dst, item2)
    }
    return dst
}


