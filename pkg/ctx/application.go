package ctx

import (
	"context"
	"gomarket/cmd"
	"gomarket/internal/errs"
)

// CtxWithApp will clone parent into a new one with the application injected on it.
func CtxWithApp(parent context.Context, app cmd.Application) context.Context {
	return context.WithValue(parent, appCtxKey, app)
}

// AppFromCtx will attempt to retrieve a cmd.Application from the current context. If it's not present or castable, it will release a panic error to be handled.
func AppFromCtx(child context.Context) cmd.Application {
	appInterface := child.Value(appCtxKey)
	if appInterface == nil {
		panic(errs.ApplicationNotPresentInContextErr)
	}

	if app, castable := appInterface.(cmd.Application); castable {
		return app
	}

	panic(errs.InvalidDataStoredInContextErr)
}
