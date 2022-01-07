package main

import (
	"flag"
	"log"
)

var (
	providers = map[string]SecretProvider{
		"gcp": &GCPSecretManager{},
	}

	templates = map[string]string{
		"plaintext": `export {{ .Name  | ToUpper }}="{{ .Data }}"`,
		"json":      `export {{ .Name  | ToUpper }}_{{ .ContentKey | ToUpper }}="{{ .ContentValue }}"`,
	}

	parsers = map[string]ContentParser{
		"plaintext": &NoParser{},
		"json":      &JSONContentParser{},
	}

	tmplLoop = `{{ range . }}
%s
{{- end -}}
`
)

// Options represents the command line options
type Options struct {
	Provider string
	Project  string
	Prefix   string
	Parser   string
	Template string
	Output   string
}

func main() {
	options := &Options{}

	flag.StringVar(&options.Provider, "provider", "gcp", "")
	flag.StringVar(&options.Project, "project", "", "")
	flag.StringVar(&options.Prefix, "secret-prefix", "", "")
	flag.StringVar(&options.Parser, "secret-data-parser", "plaintext", "")
	flag.StringVar(&options.Template, "template", "", "")
	flag.StringVar(&options.Output, "output", "", "")
	v := flag.Bool("version", false, "")

	flag.Parse()

	if *v {
		GetHumanVersion()
		return
	}

	err := Run(options)
	if err != nil {
		log.Fatal(err)
	}
}
