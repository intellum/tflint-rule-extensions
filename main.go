package main

import (
	"github.com/intellum/tflint-rule-extensions/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "google_organization_iam_binding_rule_set", // Give it a meaningful name
			Version: "0.1.0",                                    // Set the version
			Rules: []tflint.Rule{
				rules.NewGoogleOrganizationIamBindingRule(),
			},
		},
	})
}
