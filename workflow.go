package temporal_template_workflowcheck

import (
	"bytes"
	"go.temporal.io/sdk/workflow"
	"text/template"
)

var (
	Template = template.Must(template.New("sent").Parse(`Data {{.Foo}} {{.Bar}}`))
)

type TemplateInfo struct{ Foo, Bar string }

func Workflow(ctx workflow.Context) error {
	buf := new(bytes.Buffer)
	err := Template.Execute(buf, TemplateInfo{Bar: "bar", Foo: "foo"})
	if err != nil {
		return err
	}

	return nil
}
