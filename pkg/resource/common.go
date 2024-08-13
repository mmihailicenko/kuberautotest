package resource

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

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
