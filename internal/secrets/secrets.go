package secrets

import (
	"errors"
	"github.com/leocomelli/secrets-init/pkg/provider"
	"github.com/leocomelli/secrets-init/pkg/provider/common"
	"go.uber.org/zap"
	"os"
)

var (
	ErrProjectNotFound     = errors.New("project is required")
	ErrUnsupportedProvider = errors.New("unsupported provider")
	ErrUnsupportedParser   = errors.New("unsupported parser")
)

// Sync has the responsibility to synchronize external secret at container initialization
type Sync struct {
	provider provider.SecretProvider
	parser   ContentParser
	template string
	params   map[string]string
	project  string
	filter   string
	output   string
}

// Secrets return a secrets based on a set of parameters
func (s *Sync) Secrets() error {

	logger.Info("start synchronize secrets with options",
		zap.String("provider", s.provider.Name()),
		zap.String("project", s.project),
		zap.String("filter", s.filter),
		zap.String("parser", s.parser.Name()),
		zap.String("template", s.template),
		zap.String("output", s.output),
	)

	if err := s.provider.Init(s.params); err != nil {
		return err
	}

	data, err := s.provider.ListSecrets(s.project, s.filter)
	if err != nil {
		return err
	}

	common.Logger.Info("secrets found", zap.Int("len", len(data)))

	output := os.Stdin
	if s.output != "" {
		output, err = os.Create(s.output)
		if err != nil {
			return err
		}
	}

	w, err := NewWriter(output, s.template)
	if err != nil {
		return err
	}

	for _, v := range data {
		e := s.parser.Parse(v)
		err := w.Write(e...)
		if err != nil {
			return err
		}
	}

	logger.Info("finish synchronize secrets",
		zap.String("template", s.template),
		zap.String("output", s.output),
	)

	return nil
}

// New set secret options based on parameters
func New(options *common.SecretsOpts) (*Sync, error) {
	external, ok := providers[options.Provider]
	if !ok {
		return nil, ErrUnsupportedProvider
	}

	if external.Name() == "gcp" && options.Project == "" {
		return nil, ErrProjectNotFound
	}

	parser, ok := parsers[options.Parser]
	if !ok {
		return nil, ErrUnsupportedParser
	}

	template := templates[options.Parser]
	if options.Template != "" {
		template = options.Template
	}

	return &Sync{
		params: map[string]string{
			AssumeRoleKey: options.AssumeRole,
		},
		project:  options.Project,
		filter:   options.Filter,
		output:   options.Output,
		template: template,
		provider: external,
		parser:   parser,
	}, nil
}
