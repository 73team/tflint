// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule checks the pattern is valid
type AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerNotebookInstanceInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsSagemakerNotebookInstanceInvalidKmsKeyIDRule() *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule {
	return &AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule{
		resourceType:  "aws_sagemaker_notebook_instance",
		attributeName: "kms_key_id",
		max:           2048,
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Name() string {
	return "aws_sagemaker_notebook_instance_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerNotebookInstanceInvalidKmsKeyIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"kms_key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`kms_key_id does not match valid pattern ^.*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}