package main

import "text/template"

// ResourcesWriter generate code for a goa application resources.
// Resources are data structures initialized by the application handlers and passed to controller
// actions.
type ResourcesWriter struct {
	*CodeWriter
	ResourceTmpl *template.Template
}

// NewResourcesWriter returns a contexts code writer.
// Resources provide the glue between the underlying request data and the user controller.
func NewResourcesWriter(filename string) (*ResourcesWriter, error) {
	cw, err := NewCodeWriter(filename)
	if err != nil {
		return nil, err
	}
	resourceTmpl, err := template.New("resource").Funcs(cw.FuncMap).Parse(resourceT)
	if err != nil {
		return nil, err
	}
	w := ResourcesWriter{
		CodeWriter:   cw,
		ResourceTmpl: resourceTmpl,
	}
	return &w, nil
}

// Write writes the code for the context types to outdir.
func (w *ResourcesWriter) Write(targetPack string) error {
	if err := w.WriteHeader(targetPack); err != nil {
		return err
	}
	return nil
}

const (
	resourceT = `package {{.}}`
)
