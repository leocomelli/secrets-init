package main

// SecretData represents a secret thta is store in a given Secret Manager provider
type SecretData struct {
	Path         string
	Name         string
	Data         string
	ContentKey   string
	ContentValue string
}
