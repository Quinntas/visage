package sharedRouter

import (
	"context"
	"github.com/quinntas/visage/api/shared/versions"
	v1Router "github.com/quinntas/visage/api/v1/infra/router"
	"github.com/quinntas/visage/internal/api"
)

func Create() map[uint8]api.VersionRouter {
	return map[uint8]api.VersionRouter{
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
