package main

import (
	"fmt"
	"io"
	"text/template"
)

// Writer contains the actions to write the contents of secrets
type Writer struct {
	Path   string
	Writer io.Writer
	Tmpl   *template.Template
}

// NewWriter creates a new writer
func NewWriter(wr io.Writer, tmpl string) (*Writer, error) {
	t, err := template.New("").Parse(fmt.Sprintf(tmplLoop, tmpl))
	if err != nil {
		return nil, err
	}

	return &Writer{
		Writer: wr,
		Tmpl:   t,
	}, nil
}

// Write writes the secrets in a file
func (w *Writer) Write(s ...*SecretData) error {
	return w.Tmpl.Execute(w.Writer, s)
}
