package resource

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"kubeclusterautotest/pkg/test_context"
	"sigs.k8s.io/e2e-framework/klient/wait"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"time"
)

/*
Wait Functions
*/

type ReadinessChecker func(context.Context, *envconf.Config) (bool, error)

const (
	defaultTimeout      = 2 * time.Minute
	defaultPollInterval = 2 * time.Second
)

func WaitFor(tc *test_context.TestContext, readinessChecker ReadinessChecker) error {
	return wait.For(
		func(ctx context.Context) (bool, error) {
			return readinessChecker(ctx, tc.Config)
		},
		wait.WithTimeout(defaultTimeout),
		wait.WithInterval(defaultPollInterval),
	)
}

/*
Resource Builder
*/

type Option func(object runtime.Object)

type Builder interface {
	Build(options ...Option) runtime.Object
}

func NewResource(builder Builder, options ...Option) runtime.Object {
	return builder.Build(options...)
}

func WithName(name string) Option {
	return func(object runtime.Object) {
		if meta, ok := object.(metav1.Object); ok {
			meta.SetName(name)
		}
	}
}

func WithNamespace(namespace string) Option {
	return func(object runtime.Object) {
		if meta, ok := object.(metav1.Object); ok {
			meta.SetNamespace(namespace)
		}
	}
}
