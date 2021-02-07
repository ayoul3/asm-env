package asm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewFromRegion returns a session given a region
func NewFromRegion(region string) client.ConfigProvider {
	return session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
}

// NewFromRegionAndProfile returns a session given a region and aws profile
func NewFromRegionAndProfile(region string, awsProfile string) (client.ConfigProvider, error) {
	config := aws.Config{
		Region: aws.String(region),
	}
	return session.NewSessionWithOptions(session.Options{
		Profile: awsProfile,
		Config:  config,
	})
}

// New returns a new default session
func NewSession() client.ConfigProvider {
	return session.Must(session.NewSession())
}
