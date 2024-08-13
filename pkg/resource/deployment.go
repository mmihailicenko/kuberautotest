package resource

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"kubeclusterautotest/pkg/test_context"
)

/*
Deployment Object Functions
*/

func CreateDeployment(tc *test_context.TestContext, obj runtime.Object) error {
	deployment, ok := obj.(*appsv1.Deployment)
	if !ok {
		return fmt.Errorf("failed to create deployment: invalid type")
	}

	return tc.Config.Client().Resources(tc.Config.Namespace()).Create(tc.Context, deployment)
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
