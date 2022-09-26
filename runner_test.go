package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProvider struct {
	GenericProvider
}

func (m *MockProvider) Name() string {
	return "mock"
}

func (m *MockProvider) Init() error {
	return nil
}

func (m *MockProvider) ListSecrets(project string, prefix string) ([]*SecretData, error) {
	return []*SecretData{
		&SecretData{
			Path: "/project/123/secrets/mysecret",
			Name: "mysecret",
			Data: "s3cr3t",
		},
	}, nil
}

func TestUnsupportedProvider(t *testing.T) {
	opts := &Options{
		Provider: "unsupported",
		Project:  "my-project",
		Parser:   "json",
	}

	err := Run(opts)
	assert.Equal(t, ErrUnsupportedProvider, err)
}

func TestUnsupportedParser(t *testing.T) {
	opts := &Options{
		Provider: "gcp",
		Project:  "my-project",
		Parser:   "unsupported",
	}

	err := Run(opts)
	assert.Equal(t, ErrUnsupportedParser, err)
}

func TestGCPProjectNotFound(t *testing.T) {
	opts := &Options{Provider: "gcp"}
	err := Run(opts)
	assert.Equal(t, ErrProjectNotFound, err)
}

func TestHappyPath(t *testing.T) {
	providers["mock"] = &MockProvider{}

	file, err := os.CreateTemp("", "test")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	opts := &Options{
		Provider: "mock",
		Project:  "my-project",
		Parser:   "plaintext",
		Output:   file.Name(),
	}

	err = Run(opts)
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}

	expected := `
export MYSECRET="s3cr3t"`

	assert.Equal(t, expected, string(content))
}
