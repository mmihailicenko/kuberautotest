package resource

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"kubeclusterautotest/pkg/test_context"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

/*
Deployment Object Functions
*/

func CreateDeployment(tc *test_context.TestContext, deployment *appsv1.Deployment) error {
	if err := tc.Config.Client().Resources().Create(tc.Context, deployment); err != nil {
		return fmt.Errorf("failed to create deployment: %w", err)
	}
	return WaitForDeploymentReady(tc, deployment)
}

func GetDeployment(tc *test_context.TestContext, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	if err := tc.Config.Client().Resources().Get(tc.Context, deployment.GetName(), deployment.GetNamespace(), deployment); err != nil {
		return nil, fmt.Errorf("failed to get deployment: %w", err)
	}
	return deployment, nil
}

func WaitForDeploymentReady(tc *test_context.TestContext, deployment *appsv1.Deployment) error {
	readinessChecker := func(ctx context.Context, config *envconf.Config) (bool, error) {
		dep, err := GetDeployment(tc, deployment)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("failed to get deployment: %w", err)
		}
		return dep.Status.ReadyReplicas == *dep.Spec.Replicas, nil
	}
	//in progress
	if err := WaitFor(tc, readinessChecker); err != nil {
		return fmt.Errorf("deployment not ready in time: %w", err)
	}
	return nil
}

/*
Deployment Builder
*/

type DeploymentBuilder struct{}

func (db DeploymentBuilder) Build(options ...Option) runtime.Object {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "default", // Default label
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "default", // Default label
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: "default-container", // Default container name
						},
					},
				},
			},
		},
	}

	for _, option := range options {
		option(deployment)
	}

	return deployment
}

func WithReplicas(replicas int32) Option {
	return func(obj runtime.Object) {
		if deployment, ok := obj.(*appsv1.Deployment); ok {
			deployment.Spec.Replicas = &replicas
		}
	}
}

func WithContainerImage(image string) Option {
	return func(obj runtime.Object) {
		if deployment, ok := obj.(*appsv1.Deployment); ok {
			if len(deployment.Spec.Template.Spec.Containers) > 0 {
				deployment.Spec.Template.Spec.Containers[0].Image = image
			}
		}
	}
}
