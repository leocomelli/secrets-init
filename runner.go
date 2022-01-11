package main

import (
	"errors"
	"os"
)

var (
	ErrProjectNotFound     = errors.New("project is required")
	ErrUnsupportedProvider = errors.New("unsupported provider")
	ErrUnsupportedParser   = errors.New("unsupported parser")
)

func Run(options *Options) error {
	if options.Project == "" {
		return ErrProjectNotFound
	}

	parser, ok := parsers[options.Parser]
	if !ok {
		return ErrUnsupportedParser
	}

	template := templates[options.Parser]
	if options.Template != "" {
		template = templates[options.Parser]
	}

	provider, ok := providers[options.Provider]
	if !ok {
		return ErrUnsupportedProvider

	}

	if err := provider.Init(); err != nil {
		return err
	}

	data, err := provider.ListSecrets(options.Project, options.Filter)
	if err != nil {
		return err
	}

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
