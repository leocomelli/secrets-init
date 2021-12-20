package main

// SecretProvider defines the behaviors for a secret provider
type SecretProvider interface {
	Init() error
	ListSecrets(string, string) ([]*SecretData, error)
}
