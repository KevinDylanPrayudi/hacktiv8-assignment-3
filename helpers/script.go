package helpers

import "html/template"

func Script(data string) template.JS {
	return template.JS(data)
}
