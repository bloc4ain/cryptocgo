package routing

import "html/template"

var (
	tmplIndex *template.Template
)

func init() {
	var err error

	if tmplIndex, err = template.ParseFiles("webui/templates/index.gohtml"); err != nil {
		panic(err)
	}
}
