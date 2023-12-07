package common

import "go.uber.org/zap"

// SecretData represents a secret thta is store in a given Secret Manager provider
type SecretData struct {
	Path         string
	Name         string
	Data         string
	ContentKey   string
	ContentValue string
}

// SecretsOpts represents the command line options
type SecretsOpts struct {
	Provider   string
	AssumeRole string
	Project    string
	Filter     string
	Parser     string
	Template   string
	Output     string
}

const (
	AssumeRoleKey = "assume-role"
)

var (
	Logger, _ = zap.NewProduction()
)
