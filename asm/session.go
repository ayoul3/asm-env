package asm

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
)

// NewFromRegion returns a session given a region
func NewFromRegion(region string) client.ConfigProvider {

	return session.Must(session.NewSession(&aws.Config{
		Region:     aws.String(region),
		HTTPClient: getHTTPCLient(),
	}))
}

func getHTTPCLient() *http.Client {
	if os.Getenv("ASM_SKIP_SSL") == "true" {
		log.Debug("Skipping SSL cert verification")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		return &http.Client{Transport: tr}
	}
	return http.DefaultClient
}

// NewFromRegionAndProfile returns a session given a region and aws profile
func NewFromRegionAndProfile(region string, awsProfile string) (client.ConfigProvider, error) {
	config := aws.Config{
		Region:     aws.String(region),
		HTTPClient: getHTTPCLient(),
	}
	return session.NewSessionWithOptions(session.Options{
		Profile: awsProfile,
		Config:  config,
	})
}

// New returns a new default session
func NewSession() client.ConfigProvider {
	return session.Must(session.NewSession(&aws.Config{HTTPClient: getHTTPCLient()}))
}
