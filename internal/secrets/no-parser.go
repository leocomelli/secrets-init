package secrets

import (
	"github.com/leocomelli/secrets-init/pkg/provider/common"
)

// NoParser represents no parser
type NoParser struct {
	Tmpl string
}

// Name returns the strategy name
func (n *NoParser) Name() string {
	return "plaintext"
}

// Parse puts the secret data in a slice
func (n *NoParser) Parse(s *common.SecretData) []*common.SecretData {
	return []*common.SecretData{s}
}
