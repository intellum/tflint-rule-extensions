package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleOrganizationIamBindingRule detects the use of google_organization_iam_binding
type GoogleOrganizationIamBindingRule struct {
	tflint.DefaultRule
}

// NewGoogleOrganizationIamBindingRule creates a new rule instance
func NewGoogleOrganizationIamBindingRule() *GoogleOrganizationIamBindingRule {
	return &GoogleOrganizationIamBindingRule{}
}

// Name returns the rule name
func (r *GoogleOrganizationIamBindingRule) Name() string {
	return "google_organization_iam_binding"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleOrganizationIamBindingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleOrganizationIamBindingRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns a reference link for the rule
func (r *GoogleOrganizationIamBindingRule) Link() string {
	return "https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_organization_iam_binding"
}

// Check runs the logic to detect the resource
func (r *GoogleOrganizationIamBindingRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_organization_iam_binding", &hclext.BodySchema{}, nil)
	if err != nil {
		return err
	}

	logger.Debug(fmt.Sprintf("Found %d google_organization_iam_binding resources", len(resources.Blocks)))

	for _, resource := range resources.Blocks {
		err := runner.EmitIssue(
			r,
			"google_organization_iam_binding usage detected.",
			resource.DefRange, // Use DefRange to get the correct location of the resource in code
		)
		if err != nil {
			return err
		}
	}

	return nil
}
