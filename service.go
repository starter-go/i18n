package i18n

// Service ...
type Service interface {

	// Available 列出所有支持的语言，按默认的优先级顺序
	Available() []Language

	// Default 获取默认的资源链
	Default() Resources

	// GetResources 按照指定的顺序，获取对应的资源链; 越靠前(index 越小)，优先级越高
	GetResources(l ...Language) Resources
}
