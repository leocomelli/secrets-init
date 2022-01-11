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
	Filter   string
	Parser   string
	Template string
	Output   string
}

func main() {
	options := &Options{}

	flag.StringVar(&options.Provider, "provider", "gcp", "name of the provider that manages the secrets")
	flag.StringVar(&options.Project, "project", "", "gcp project that contains the secrets")
	flag.StringVar(&options.Filter, "filter", "", "regex to filter secrets by name")
	flag.StringVar(&options.Parser, "data-parser", "plaintext", "parse secret based on data type")
	flag.StringVar(&options.Template, "template", "", "template to render secret data")
	flag.StringVar(&options.Output, "output", "", "path to write output file to")
	v := flag.Bool("version", false, "show the current secrets-init version")

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
