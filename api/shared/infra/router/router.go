package sharedRouter

import (
	"context"
	"github.com/quinntas/visage/api/shared/versions"
	v1Router "github.com/quinntas/visage/api/v1/infra/router"
	"github.com/quinntas/visage/internal/api"
)

func Create() api.RouterMap {
	return api.RouterMap{
		versions.V1: v1Router.Create(),
	}
}

func CreateWithContext(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		api.ROUTERS_CTX_KEY,
		Create(),
	)
}
