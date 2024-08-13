package deployment

import (
	"kubeclusterautotest/pkg/test_context"
	"sigs.k8s.io/e2e-framework/pkg/features"
	"testing"
)

func TestDeployment(t *testing.T) {
	testDeployment := features.NewWithDescription("Deployment Test",
		"Basic Deployment Test Scenario. Deployment is created, assessed and deleted").
		WithSetup("1. Deployment Setup", test_context.WithTestContext(SetupBasic)).
		Assess("2. Deployment Assessment", test_context.WithTestContext(AssessDeployment)).
		WithTeardown("3. Deployment Deletion", test_context.WithTestContext(DeleteDeployment)).
		Feature()
	testEnv.Test(t, testDeployment)
}
