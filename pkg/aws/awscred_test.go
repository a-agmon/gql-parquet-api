package aws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAWSCredEnv(t *testing.T) {
	// Set up test environment variables
	os.Setenv("AWS_ACCESS_KEY_ID", "access_key")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret_key")
	os.Setenv("AWS_SESSION_TOKEN", "session_token")
	os.Setenv("AWS_REGION", "us-west-2")

	// Call the function being tested
	cred, err := GetAWSCredEnv()

	// Check the results
	if err != nil {
		t.Errorf("GetAWSCredEnv returned an error: %v", err)
	}
	if cred.AccessKeyID != "access_key" {
		t.Errorf("AccessKeyID is incorrect: got %v, want %v", cred.AccessKeyID, "access_key")
	}

	os.Unsetenv("AWS_REGION")

	_, err = GetAWSCredEnv()
	assert.Error(t, err)

	// Clean up test environment variables
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")
	os.Unsetenv("AWS_SESSION_TOKEN")

}
