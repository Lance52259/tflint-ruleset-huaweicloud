// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleBigqueryRoutineInvalidLanguageRule checks the pattern is valid
type GoogleBigqueryRoutineInvalidLanguageRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleBigqueryRoutineInvalidLanguageRule returns new rule with default attributes
func NewGoogleBigqueryRoutineInvalidLanguageRule() *GoogleBigqueryRoutineInvalidLanguageRule {
	return &GoogleBigqueryRoutineInvalidLanguageRule{
		resourceType:  "google_bigquery_routine",
		attributeName: "language",
	}
}

// Name returns the rule name
func (r *GoogleBigqueryRoutineInvalidLanguageRule) Name() string {
	return "google_bigquery_routine_invalid_language"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleBigqueryRoutineInvalidLanguageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleBigqueryRoutineInvalidLanguageRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleBigqueryRoutineInvalidLanguageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleBigqueryRoutineInvalidLanguageRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"SQL", "JAVASCRIPT", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}