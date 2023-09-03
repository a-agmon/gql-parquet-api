package aws

import (
	"errors"
	"os"
)

type AWSCred struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Region          string
}

func GetAWSCredEnv() (AWSCred, error) {
	c := AWSCred{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
		Region:          os.Getenv("AWS_REGION"),
	}
	if c.AccessKeyID == "" || c.SecretAccessKey == "" || c.SessionToken == "" || c.Region == "" {
		return c, errors.New("one or more of the following env vars are not set: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN, AWS_REGION")
	}
	return c, nil
}
