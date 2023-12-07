package provider

import (
	"github.com/leocomelli/secrets-init/pkg/provider/common"
	"go.uber.org/zap"
	"regexp"
)

var (
	logger, _ = zap.NewProduction()
)

// SecretProvider defines the behaviors for a secret provider
type SecretProvider interface {
	Name() string
	Init(map[string]string) error
	Filter(string, string) bool
	ListSecrets(string, string) ([]*common.SecretData, error)
}

type GenericProvider struct{}

// Filter the secrets by regex
func (s *GenericProvider) Filter(name string, exp string) bool {
	re := regexp.MustCompile(exp)
	return re.MatchString(name)
}
