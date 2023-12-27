package ctx_test

import (
	"context"
	"gomarket/cmd"
	"gomarket/internal/errs"
	"gomarket/pkg/ctx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CtxWithApp(t *testing.T) {
	// TODO: mockery
	// Arrange
	app := cmd.NewApp()
	parent := context.Background()

	// Act
	child := ctx.CtxWithApp(parent, app)

	// Assert
	assert.NotEqual(t, parent, child, "child context must be different than parent context")
}

func Test_AppFromCtx(t *testing.T) {
	// TODO: mockery
	app := cmd.NewApp()
	parent := context.Background()
	child := ctx.CtxWithApp(parent, app)

	t.Run("should return the same app from context", func(t *testing.T) {
		// Act
		appFromCtx := ctx.AppFromCtx(child)

		// Assert
		assert.Equal(t, app, appFromCtx, "the application extracted from context must be the same as the injected")
	})

	t.Run("should panic if app is not present in context", func(t *testing.T) {
		assert.PanicsWithError(t, errs.ApplicationNotPresentInContextErr.Error(), func() {
			ctx.AppFromCtx(parent)
		}, "the absence of the application on context must cause panic")
	})

	t.Run("should panic if data present in context is not castable as an Application", func(t *testing.T) {
		// Arrange
		const brokenAppCtxKey ctx.CtxKeys = iota
		altChild := context.WithValue(parent, brokenAppCtxKey, "not-app")

		// Act + Assert
		assert.PanicsWithError(t, errs.InvalidDataStoredInContextErr.Error(), func() {
			ctx.AppFromCtx(altChild)
		}, "the uncastable data presence in context must cause panic")
	})
}
