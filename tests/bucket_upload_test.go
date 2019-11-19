package test

import (
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/gcp"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"time"
	"crypto/tls"
)

func TestBucketUploadFileExistsTest(t *testing.T) {

	projectID := gcp.GetGoogleProjectIDFromEnvVar(t)

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform/modules/bucket-upload",
		Vars: map[string]interface{}{
			"project_id": projectID,
		},
	}

	fmt.Printf("Starting test. ProjectID = %s \n", projectID)
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	filePath := terraform.OutputRequired(t, terraformOptions, "file_path")
	// File URL should be https://storage.googleapis.com/terratest-259517_terratest/test.txt
	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s", filePath)
	fmt.Printf("FileURL = %v\n", fileURL)

	tlsConfig := tls.Config{}
	maxRetries := 5
	timeBetweenRetries := 5 * time.Second

	expectedStatus := 200
	expectedBody := "It works"
	http_helper.HttpGetWithRetry(t, fileURL, &tlsConfig, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
}
