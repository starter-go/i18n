package i18n

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName     = "github.com/starter-go/i18n"
	theModuleVersion  = "v0.0.3"
	theModuleRevision = 3
)

////////////////////////////////////////////////////////////////////////////////

const (
	theMainModuleResPath = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewMainModule 用于创建模块 ['github.com/starter-go/i18n']
func NewMainModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	mb.Depend(starter.Module())
	return mb
}

// NewTestModule 用于创建模块 ['github.com/starter-go/i18n']
func NewTestModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}
