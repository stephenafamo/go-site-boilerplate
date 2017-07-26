package config

import (
	"html/template"
)

// compile all templates and cache them
var Templates = template.Must(template.ParseGlob("/go/src/github.com/stephenafamo/site/templates/*"))
