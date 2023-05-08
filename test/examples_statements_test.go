package test

import (
	"encoding/json"
	"os"
	"testing"

	TTAWS "github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

type Statement struct {
	Sid      string   `json:"Sid"`
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource []string `json:"Resource"`
}

type Policy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

func TestTerraformStatements(t *testing.T) {
	roleName := "terraform-test-iam-policy-" + GetShortId()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/statements",
		Vars: map[string]interface{}{
			"name":    roleName,
			"account": os.Getenv("AWS_TESTING_ACCOUNT"),
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	name := terraform.Output(t, terraformOptions, "name")
	assert.NotEmpty(t, name)

	arn := terraform.Output(t, terraformOptions, "arn")
	assert.NotEmpty(t, arn)

	id := terraform.Output(t, terraformOptions, "id")
	assert.NotEmpty(t, id)

	jsonOutput := terraform.Output(t, terraformOptions, "json")
	assert.NotEmpty(t, jsonOutput)

	sess, err := NewSession(os.Getenv("AWS_REGION"))
	assert.NoError(t, err)

	client := iam.New(sess)

	input := iam.ListPoliciesInput{}
	output, err := client.ListPolicies(&input)
	assert.NoError(t, err)

	role := FilterPolicies(t, name, output.Policies)

	assert.Equal(t, "/", aws.StringValue(role.Path))

	var policy Policy
	err = json.Unmarshal([]byte(jsonOutput), &policy)
	assert.NoError(t, err)

	assert.Equal(t, "2012-10-17", policy.Version)

	assert.Equal(t, 1, len(policy.Statement))
	assert.Equal(t, "", policy.Statement[0].Sid)
	assert.Equal(t, "Allow", policy.Statement[0].Effect)
	assert.Equal(t, 5, len(policy.Statement[0].Action))
	assert.Equal(t, 2, len(policy.Statement[0].Resource))
}

func FilterPolicies(t *testing.T, name string, roles []*iam.Policy) *iam.Policy {
	for _, role := range roles {
		if aws.StringValue(role.PolicyName) == name {
			return role
		}
	}

	t.Fatalf("Could not file policy %s", name)

	return nil
}

func NewSession(region string) (*session.Session, error) {
	sess, err := TTAWS.NewAuthenticatedSession(region)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func GetShortId() string {
	githubSha := os.Getenv("GITHUB_SHA")
	if len(githubSha) >= 7 {
		return githubSha[0:6]
	}

	return "local"
}
