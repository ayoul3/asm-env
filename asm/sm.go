package asm

import (
	"os"
	"regexp"
	"strings"

	"emperror.dev/errors"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	log "github.com/sirupsen/logrus"
)

const SMPatern = "arn:aws:secretsmanager:"

// Client is a SM custom client
type Client struct {
	api    secretsmanageriface.SecretsManagerAPI
	region string
}

// NewClient returns a new Client from an AWS SM client
func NewClient(api secretsmanageriface.SecretsManagerAPI) *Client {
	region := os.Getenv("AWS_REGION")
	return &Client{
		api,
		region,
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
	secretName := c.ExtractPath(key)
	secretRegion := c.ExtractRegion(key)
	api := c.api
	if secretRegion != c.region {
		log.Debugf("Switching regions to %s for key %s", secretRegion, key)
		api = NewAPIForRegion(secretRegion)
	}
	res, err := api.GetSecretValue(new(secretsmanager.GetSecretValueInput).SetSecretId(secretName))
	if err != nil {
		return "", errors.Wrapf(err, "GetSecretValue ")
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
	var re = regexp.MustCompile(`arn:aws:secretsmanager:([a-z0-9-]+):\d+:`)
	match := re.FindStringSubmatch(key)
	if len(match) < 1 {
		return c.region
	}
	return match[1]
}
