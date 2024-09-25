package entsetorclear

import (
	_ "embed"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

//go:embed setorclear.tpl
var tmplfile string

// Extension implements entc.Extension.
type Extension struct {
	entc.DefaultExtension
}

func (Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("entsetorclear").Parse(tmplfile)),
	}
}

func New() *Extension {
	return &Extension{}
}
