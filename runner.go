package main

import (
	"errors"
	"os"

	"go.uber.org/zap"
)

var (
	ErrProjectNotFound     = errors.New("project is required")
	ErrUnsupportedProvider = errors.New("unsupported provider")
	ErrUnsupportedParser   = errors.New("unsupported parser")
)

func Run(options *Options) error {
	provider, ok := providers[options.Provider]
	if !ok {
		return ErrUnsupportedProvider
	}

	if provider.Name() == "gcp" && options.Project == "" {
		return ErrProjectNotFound
	}

	parser, ok := parsers[options.Parser]
	if !ok {
		return ErrUnsupportedParser
	}

	template := templates[options.Parser]
	if options.Template != "" {
		template = options.Template
	}

	logger.Info("using options", zap.Any("values", options))

	params := map[string]string{
		AssumeRoleKey: options.AssumeRole,
	}

	if err := provider.Init(params); err != nil {
		return err
	}

	data, err := provider.ListSecrets(options.Project, options.Filter)
	if err != nil {
		return err
	}

	logger.Info("secrets found", zap.Int("len", len(data)))

	output := os.Stdin
	if options.Output != "" {
		output, err = os.Create(options.Output)
		if err != nil {
			return err
		}
	}

	w, err := NewWriter(output, template)
	if err != nil {
		return err
	}

	for _, v := range data {
		e := parser.Parse(v)
		err := w.Write(e...)
		if err != nil {
			return err
		}
	}

	return nil
}
