package test_context

import (
	"context"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
	"testing"
)

type TestContext struct {
	Context context.Context
	T       *testing.T
	Config  *envconf.Config
}

func NewTestContext(ctx context.Context, t *testing.T, cfg *envconf.Config) *TestContext {
	return &TestContext{Context: ctx, T: t, Config: cfg}
}

func WithTestContext(f func(*TestContext) error) features.Func {
	return func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		tc, ok := ctx.Value("testContext").(*TestContext)
		if !ok {
			t.Logf("New test context is created")
			tc = NewTestContext(ctx, t, cfg)
			ctx = context.WithValue(ctx, "testContext", tc)
		}
		if err := f(tc); err != nil {
			t.Fatal("Fail to create new test context", err)
		}
		return ctx
	}
}
