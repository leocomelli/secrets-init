package secrets

import (
	"github.com/leocomelli/secrets-init/pkg/provider"
	"github.com/leocomelli/secrets-init/pkg/provider/common"
	"go.uber.org/zap"
)

// ContentParser defines secret content parser behaviors
type ContentParser interface {
	Parse(*common.SecretData) []*common.SecretData
	Name() string
}

var (
	logger, _ = zap.NewProduction()

	providers = map[string]provider.SecretProvider{
		"gcp": &provider.GCPSecretManager{},
		"aws": &provider.AWSSecretManager{},
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

const (
	AssumeRoleKey = "assume-role"
)
