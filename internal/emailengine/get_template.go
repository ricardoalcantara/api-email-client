package emailengine

import (
	"bytes"
	"text/template"
)

func GetTemplate(templateStr string, context any) string {
	t := template.Must(template.New("tmp").Parse(templateStr))

	var body bytes.Buffer
	err := t.Execute(&body, context)
	if err != nil {
		panic(err)
	}

	return body.String()
}
