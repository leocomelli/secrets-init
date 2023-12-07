package provider

import (
	"context"
	"github.com/leocomelli/secrets-init/pkg/provider/common"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	stsTypes "github.com/aws/aws-sdk-go-v2/service/sts/types"

	"go.uber.org/zap"
	"math/rand"
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
func (s *AWSSecretManager) Init(params map[string]string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	if role, ok := params[common.AssumeRoleKey]; ok && role != "" {
		cfg, err = s.AssumeRole(cfg, role)
		if err != nil {
			return err
		}
	}

	svc := secretsmanager.NewFromConfig(cfg)
	s.client = svc

	return nil
}

func (s *AWSSecretManager) AssumeRole(cfg aws.Config, role string) (aws.Config, error) {
	sourceAccount := sts.NewFromConfig(cfg)

	logger.Info("Assuming role", zap.String("role", role))

	rand.Seed(time.Now().UnixNano())
	response, err := sourceAccount.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
		RoleArn:         aws.String(role),
		RoleSessionName: aws.String("sc_" + strconv.Itoa(10000+rand.Intn(25000))),
	})

	if err != nil {
		return aws.Config{}, err
	}

	var assumedRoleCreds *stsTypes.Credentials = response.Credentials

	return config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(
			*assumedRoleCreds.AccessKeyId,
			*assumedRoleCreds.SecretAccessKey,
			*assumedRoleCreds.SessionToken,
		)))
}

// ListSecrets lists the AWS Secrets using external configurations.
// Use prefix to filter the secrets starting with a term.
// If prefix is empty, all secrets are listed.
func (s *AWSSecretManager) ListSecrets(_ string, prefix string) ([]*common.SecretData, error) {
	var (
		data  []*common.SecretData
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

				data = append(data, &common.SecretData{
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
