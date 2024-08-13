package deployment

import (
	appsv1 "k8s.io/api/apps/v1"
	"kubeclusterautotest/pkg/resource"
	"kubeclusterautotest/pkg/test_context"
	"os"
	"testing"

	"sigs.k8s.io/e2e-framework/klient/conf"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/envfuncs"
)

var testEnv env.Environment

func TestMain(m *testing.M) {
	testEnv = env.New()
	path := conf.ResolveKubeConfigFile()
	cfg := envconf.NewWithKubeConfig(path)
	testEnv = env.NewWithConfig(cfg)
	namespace := envconf.RandomName("sample-ns", 16)

	testEnv.Setup(
		envfuncs.CreateNamespace(namespace),
	)

	testEnv.Finish(
		envfuncs.DeleteNamespace(namespace),
	)

	os.Exit(testEnv.Run(m))
}

func SetupBasic(tc *test_context.TestContext) error {
	basicDeployment := resource.NewResource(resource.DeploymentBuilder{},
		resource.WithName("test-deployment"),
		resource.WithNamespace(tc.Config.Namespace()),
		resource.WithReplicas(1),
		resource.WithContainerImage("nginx"),
	).(*appsv1.Deployment)
	if err := resource.CreateDeployment(tc, basicDeployment); err != nil {
		return err
	}
	return nil
}

func AssessDeployment(tc *test_context.TestContext) error {
	//assessment logic
	return nil
}

func DeleteDeployment(tc *test_context.TestContext) error {
	//deletion logic
	return nil
}
