package asm

import (
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/prometheus/common/log"
)

const SMPatern = "arn:aws:secretsmanager:"

// Client is a SM custom client
type Client struct {
	api secretsmanageriface.SecretsManagerAPI
}

// NewClient returns a new Client from an AWS SM client
func NewClient(api secretsmanageriface.SecretsManagerAPI) *Client {
	return &Client{
		api,
	}
}

// NewAPI returns a new concrete AWS SM client
func NewAPI() *secretsmanager.SecretsManager {
	return secretsmanager.New(NewSession())
}

// NewAPIForRegion returns a new concrete AWS SM client for a specific region
func NewAPIForRegion(region string) secretsmanageriface.SecretsManagerAPI {
	return secretsmanager.New(NewFromRegion(region))
}

// GetSecret return a Secret fetched from SM
func (c *Client) GetSecret(key string) (secret string, err error) {
	formattedKey := c.ExtractPath(key)
	res, err := c.api.GetSecretValue(new(secretsmanager.GetSecretValueInput).SetSecretId(formattedKey))
	if err != nil {
		return "", err
	}
	return *res.SecretString, nil
}

// Overrides region
func (c *Client) WithRegion(region string) {
	c.api = secretsmanager.New(NewFromRegion(region))
}

func (c *Client) IsSecret(key string) bool {
	return strings.Contains(key, "arn:aws:secretsmanager")
}

func (c *Client) ExtractPath(key string) (out string) {
	var re = regexp.MustCompile(`(arn:aws:secretsmanager:[a-z0-9-]+:\d+:secret:[a-zA-Z0-9/._=@-]+)`)
	match := re.FindStringSubmatch(key)
	if len(match) < 1 {
		log.Warnf("Badly formatted key %s", key)
		return key
	}
	return match[1]
}

func (c *Client) ExtractRegion(key string) (region string) {
	var re = regexp.MustCompile(`(arn:aws:secretsmanager:([a-z0-9-]+):\d+:`)
	match := re.FindStringSubmatch(key)
	if len(match) < 1 {
		return "eu-west-1"
	}
	return match[1]
}
