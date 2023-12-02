package i18n

// Service ...
type Service interface {
	Available() []Language
	GetResources(l ...Language) Resources
	Default() Resources
}
