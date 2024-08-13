package test

import (
	"kubeclusterautotest/pkg/resource"
	"kubeclusterautotest/pkg/test_context"
	"sigs.k8s.io/e2e-framework/pkg/features"
	"testing"
)

func TestDeployment(t *testing.T) {
	depTest := features.NewWithDescription("Deployment Test",
		"1. Deployment setup, "+
			"2. Deployment assessment, "+
			"3. Deployment Teardown").
		WithSetup("1. Deployment setup", test_context.WithTestContext(func(tc *test_context.TestContext) error {
			deployment := resource.NewResource(resource.DeploymentBuilder{},
				resource.WithName("test-deployment"),
				resource.WithNamespace(tc.Config.Namespace()),
				resource.WithReplicas(3),
				resource.WithContainerImage("nginx"),
			)
			return resource.CreateDeployment(tc, deployment)
		})).Feature()
	testEnv.Test(t, depTest)
}
