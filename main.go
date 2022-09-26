package main

import (
	"flag"
	"fmt"

	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()

	providers = map[string]SecretProvider{
		"gcp": &GCPSecretManager{},
		"aws": &AWSSecretManager{},
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

func (o *Options) String() string {
	return fmt.Sprintf(
		"provider: %s, project: %s, filter: %s, parser: %s, template: %s, output: %s",
		o.Provider, o.Project, o.Filter, o.Parser, o.Template, o.Output,
	)
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
		logger.Fatal("error getting the secrets", zap.Error(err))
	}
}
