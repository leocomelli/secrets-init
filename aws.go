package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"go.uber.org/zap"
)

// AWSSecretManager represents the AWS Secret Manager
type AWSSecretManager struct {
	GenericProvider
	client *secretsmanager.Client
}

// Name returns the provider name
func (s *AWSSecretManager) Name() string {
	return "aws"
}

// Init initializes the AWSSecretManager instance that
// contains a client for interacting with Secret Manager API.
//
// It populates an AWS Config with the values from the external
// configurations.
//
// The default configuration sources are:
// * Environment Variables
// * Shared Configuration and Shared Credentials files.
func (s *AWSSecretManager) Init() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	svc := secretsmanager.NewFromConfig(cfg)
	s.client = svc

	return nil
}

// ListSecrets lists the AWS Secrets using external configurations.
// Use prefix to filter the secrets starting with a term.
// If prefix is empty, all secrets are listed.
func (s *AWSSecretManager) ListSecrets(_ string, prefix string) ([]*SecretData, error) {
	var (
		data  []*SecretData
		token string
	)

	input := &secretsmanager.ListSecretsInput{}

	for {
		result, err := s.client.ListSecrets(context.Background(), input)

		if err != nil {
			return nil, err
		}

		token = aws.ToString(result.NextToken)

		for _, secret := range result.SecretList {
			name := aws.ToString(secret.Name)

			if s.Filter(name, prefix) {
				resp, err := s.client.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{
					SecretId: secret.Name,
				})

				if err != nil {
					logger.Error("error getting the secret value", zap.String("name", name))
					continue
				}

				data = append(data, &SecretData{
					Path: name,
					Name: name,
					Data: aws.ToString(resp.SecretString),
				})
			}
		}

		if token != "" {
			input.NextToken = aws.String(token)
		} else {
			break
		}
	}

	return data, nil
}
