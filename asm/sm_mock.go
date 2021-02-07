package asm

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// MockSecretValue is a dummy value for the mocks
const MockSecretValue string = "@MY_SECRET_VALUE@"

// MockClient is an AWS secretsmanager client mock
type MockClient struct {
	secretsmanageriface.SecretsManagerAPI
	GetSecretShouldFail      bool
	DescribeSecretShouldFail bool
	ShouldBeEmpty            bool
	SecretValue              string
}

// GetSecretValue is a mock implementation of secretsmanager method
func (m *MockClient) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	if m.GetSecretShouldFail {
		return nil, errors.New("GetSecretValue was forced to fail")
	}
	secret := MockSecretValue
	if m.SecretValue != "" {
		secret = m.SecretValue
	}
	output := new(secretsmanager.GetSecretValueOutput).SetSecretString(secret)
	return output, nil
}

// DescribeSecret is a mock implementation of the sm method
func (m *MockClient) DescribeSecret(input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error) {
	if m.DescribeSecretShouldFail {
		return nil, errors.New("DescribeSecret was forced to fail")
	}
	secret := new(secretsmanager.DescribeSecretOutput)
	if !m.ShouldBeEmpty {
		secret.SetTags([]*secretsmanager.Tag{
			{
				Key:   aws.String("MY_KEY"),
				Value: aws.String("MY_VALUE"),
			},
		})
	}

	return secret, nil
}
