package main

import "regexp"

// SecretProvider defines the behaviors for a secret provider
type SecretProvider interface {
	Name() string
	Init(map[string]string) error
	Filter(string, string) bool
	ListSecrets(string, string) ([]*SecretData, error)
}

type GenericProvider struct{}

// Filter the secrets by regex
func (s *GenericProvider) Filter(name string, exp string) bool {
	re := regexp.MustCompile(exp)
	return re.MatchString(name)
}
