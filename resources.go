package i18n

// Resources 表示一个资源的集合
type Resources interface {
	ReadText(path string) (string, error)

	ReadBinary(path string) ([]byte, error)

	ReadJSON(path string, root any) error

	GetString(name string) (string, error)

	String(name string) string

	Names() []string
}
