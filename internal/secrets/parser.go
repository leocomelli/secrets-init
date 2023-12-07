package secrets

import (
	"encoding/json"
	"fmt"
	"github.com/leocomelli/secrets-init/pkg/provider/common"
	"go.uber.org/zap"
)

// JSONContenParser represents a JSON parser
type JSONContentParser struct {
	Tmpl string
}

// Name returns the strategy name
func (j *JSONContentParser) Name() string {
	return "json"
}

/*
Parse lists each json entry as a secret
Consider a secret called myscret with the content below

	{
	   "user": "root",
	   "password": "s3cr3t",
	   "host" : "127.0.0.1:5432",
	}

# The following secrets will be returned

	myscret_user: root
	mysecret_password: s3cr3t
	mysecret_host: 127.0.0.1:5432
*/
func (j *JSONContentParser) Parse(s *common.SecretData) []*common.SecretData {
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(s.Data), &m); err != nil {
		logger.Warn("invalid json", zap.String("name", s.Name))
	}

	secrets := make([]*common.SecretData, 0, len(m))
	for k, v := range m {
		secrets = append(secrets, &common.SecretData{
			Name:         s.Name,
			Path:         s.Path,
			Data:         s.Data,
			ContentKey:   k,
			ContentValue: fmt.Sprintf("%v", v),
		})
	}

	return secrets
}
