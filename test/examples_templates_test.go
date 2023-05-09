package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

var ErrUnsupportedType = errors.New("unsupported type")

type StrArray []string

type StatementTemplates struct {
	Sid      string   `json:"Sid"`
	Effect   string   `json:"Effect"`
	Action   StrArray `json:"Action"`
	Resource StrArray `json:"Resource"`
}

type PolicyTemplates struct {
	Version   string               `json:"Version"`
	Statement []StatementTemplates `json:"Statement"`
}

func TestTerraformTemplates(t *testing.T) {
	roleName := "terraform-test-iam-policy-templates-" + GetShortId()
	region := os.Getenv("AWS_REGION")
	accountId := os.Getenv("AWS_TESTING_ACCOUNT_ID")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/templates",
		Vars: map[string]interface{}{
			"name":    roleName,
			"account": fmt.Sprintf("arn:aws:iam::%s:root", accountId),
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

	var policy PolicyTemplates
	err = json.Unmarshal([]byte(jsonOutput), &policy)
	assert.NoError(t, err)

	assert.Equal(t, "2012-10-17", policy.Version)
	assert.Equal(t, 5, len(policy.Statement))

	AssertStatement(
		t,
		"RegisterTaskDefinition",
		"Allow",
		[]string{"*"},
		[]string{"ecs:RegisterTaskDefinition", "ecs:DescribeTaskDefinition"},
		policy.Statement,
	)

	AssertStatement(
		t,
		"PassRolesInTaskDefinition",
		"Allow",
		[]string{
			fmt.Sprintf("arn:aws:iam::%s:role/web-app", accountId),
			fmt.Sprintf("arn:aws:iam::%s:role/web-app-exec", accountId),
		},
		[]string{"iam:PassRole"},
		policy.Statement,
	)

	AssertStatement(
		t,
		"ECSDeployment",
		"Allow",
		[]string{
			fmt.Sprintf("arn:aws:ecs:%s:%s:service/my-project/web-app", region, accountId),
			fmt.Sprintf("arn:aws:codedeploy:%s:%s:deploymentgroup:my-project/web-app", region, accountId),
			fmt.Sprintf("arn:aws:codedeploy:%s:%s:deploymentconfig:*", region, accountId),
			fmt.Sprintf("arn:aws:codedeploy:%s:%s:application:my-project", region, accountId),
		},
		[]string{
			"ecs:UpdateService",
			"ecs:DescribeServices",
			"codedeploy:RegisterApplicationRevision",
			"codedeploy:GetDeploymentGroup",
			"codedeploy:GetDeploymentConfig",
			"codedeploy:GetDeployment",
			"codedeploy:CreateDeployment",
			"codedeploy:ContinueDeployment",
			"codedeploy:StopDeployment",
		},
		policy.Statement,
	)

	AssertStatement(
		t,
		"ECRAuthorization",
		"Allow",
		[]string{"*"},
		[]string{"ecr:GetAuthorizationToken"},
		policy.Statement,
	)

	AssertStatement(
		t,
		"ECRPushAndPull",
		"Allow",
		[]string{fmt.Sprintf("arn:aws:ecr:%s:%s:repository/web-app", region, accountId)},
		[]string{
			"ecr:UploadLayerPart",
			"ecr:PutImage",
			"ecr:InitiateLayerUpload",
			"ecr:GetDownloadUrlForLayer",
			"ecr:DescribeImageScanFindings",
			"ecr:CompleteLayerUpload",
			"ecr:BatchGetImage",
			"ecr:BatchCheckLayerAvailability",
		},
		policy.Statement,
	)
}

func AssertStatement(
	t *testing.T,
	sid string,
	effect string,
	resource []string,
	action []string,
	statements []StatementTemplates,
) {
	for _, statement := range statements {
		sidMatches := sid == statement.Sid
		effectMatches := effect == statement.Effect
		resourceMatches := 0 == len(difference(resource, statement.Resource))
		actionMatches := 0 == len(difference(action, statement.Action))

		if sidMatches && effectMatches && resourceMatches && actionMatches {
			return
		}
	}

	t.Errorf(
		"Could not find matching statement! sid: %s, effect: %s, resource: %v, action: %v",
		sid, effect, resource, action,
	)
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// custom json unmarshal to normalize string or string array to string array
func (sa *StrArray) UnmarshalJSON(data []byte) error {
	var jsonObj interface{}
	err := json.Unmarshal(data, &jsonObj)
	if err != nil {
		return err
	}
	switch obj := jsonObj.(type) {
	case string:
		*sa = []string{obj}
		return nil
	case []interface{}:
		s := make([]string, 0, len(obj))
		for _, v := range obj {
			value, ok := v.(string)
			if !ok {
				return ErrUnsupportedType
			}
			s = append(s, value)
		}
		*sa = s
		return nil
	}
	return ErrUnsupportedType
}
