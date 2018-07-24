package config

import "html/template"

var TPL *template.Template

func init() {
	TPL = temaplate.Must(template.ParseGlob("templates/*"))
}
