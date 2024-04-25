package v1Router

import (
	"context"
	"github.com/quinntas/visage/api/shared/implementations"
	"github.com/quinntas/visage/api/shared/versions"
	v1Publish "github.com/quinntas/visage/api/v1/implementations/publish"
	v1Subscribe "github.com/quinntas/visage/api/v1/implementations/subscribe"
	"github.com/quinntas/visage/internal/api"
)

// TODO: research how to do better ctx applies
func Create() api.VersionRouter {
	ctx := v1Subscribe.ApplySubscribeContext(context.Background())

	return api.NewVersionRouter(
		versions.V1,
		api.ImplMap{
			implementations.SUBSCRIBE: v1Subscribe.NewSubscribe(ctx),
			implementations.PUBLISH:   v1Publish.NewPublish(ctx),
		},
	)
}
