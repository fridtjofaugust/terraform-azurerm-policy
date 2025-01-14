package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUT_Examples(t *testing.T) {
	t.Parallel()

	tests := []string{
		//"../examples/custom",
		//"../examples/existing",
		//"../examples/management-group",
		//"../examples/policy-set-definitions",
		"../examples/multiple-assignments",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			tfOptions := &terraform.Options{
				TerraformDir: test,
			}

			terraform.Init(t, tfOptions)
			terraform.RunTerraformCommand(t, tfOptions, terraform.FormatArgs(tfOptions, "plan")...)
		})
	}
}
