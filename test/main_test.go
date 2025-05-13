package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/gruntwork-io/terratest/modules/random"
    "fmt"
)

func TestTerraformModule(t *testing.T) {
    t.Parallel()

    // Unique test name
    testName := "terraform-test-" + random.UniqueId()

    fmt.Println("Running test:", testName)

    terraformOptions := &terraform.Options{
        TerraformDir: "../", // Pointing to the parent folder where the module resides

        Vars: map[string]interface{}{
            "ami_id": "ami-0f88e80871fd81e91", // Replace with a valid AMI ID
            "instance_type": "t2.micro",
        },

        NoColor: true,
    }

    // Clean up after test
    defer terraform.Destroy(t, terraformOptions)

    // Initialize and apply the Terraform configuration
    terraform.InitAndApply(t, terraformOptions)

    // Validate your code works as expected
    instanceID := terraform.Output(t, terraformOptions, "instance_id")
    if instanceID == "" {
        t.Fatal("Instance ID is empty!")
    }
}
