package i18n

import "strings"

// Language 表示语言的字符串（例如：zh_cn, en_us ）
type Language string

func (lang Language) String() string {
	return string(lang)
}

// Normalize 标准化
func (lang Language) Normalize() Language {
	str := lang.String()
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	return Language(str)
}

////////////////////////////////////////////////////////////////////////////////

// LanguageRegistration ...
type LanguageRegistration struct {
	Language  Language
	Enabled   bool
	Priority  int
	Resources Resources
}

// LanguageProvider ...
type LanguageProvider interface {
	Languages() []*LanguageRegistration
}
