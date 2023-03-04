package main

// nolint:staticcheck
import (
	"context"
	"fmt"
	"path"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/iterator"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// GCPSecretManager represents the Google Cloud Platform Secret Manager
type GCPSecretManager struct {
	GenericProvider
	client *secretmanager.Client
}

// Name returns the provider name
func (s *GCPSecretManager) Name() string {
	return "gcp"
}

// Init initializes the GCPSecretManager instance that
// contains a client for interacting with Secret Manager API.
//
// It uses a local credential based on the GOOGLE_APPLICATION_CREDENTIALS
// environment variable.
//
// See: https://cloud.google.com/docs/authentication/getting-started
func (s *GCPSecretManager) Init(_ map[string]string) error {
	c, err := secretmanager.NewClient(context.Background())
	if err != nil {
		return err
	}

	s.client = c

	return nil
}

// ListSecrets lists the GCP Secrets for a given project.
// Use prefix to filter the secrets starting with a term.
// If prefix is empty, all secrets are listed.
// nolint:staticcheck
func (s *GCPSecretManager) ListSecrets(project string, prefix string) ([]*SecretData, error) {
	req := &secretmanagerpb.ListSecretsRequest{
		Parent: fmt.Sprintf("projects/%s", project),
	}
	it := s.client.ListSecrets(context.Background(), req)

	var data []*SecretData

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		name := path.Base(resp.GetName())

		if s.Filter(name, prefix) {
			content, err := s.client.AccessSecretVersion(context.Background(), &secretmanagerpb.AccessSecretVersionRequest{Name: resp.GetName() + "/versions/latest"})
			if err != nil {
				return nil, err
			}

			data = append(data, &SecretData{
				Path: resp.GetName(),
				Name: name,
				Data: string(content.Payload.Data),
			})
		}
	}

	return data, nil
}
