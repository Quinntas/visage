package v1Router

import (
	"github.com/quinntas/visage/api/shared/implementations"
	"github.com/quinntas/visage/api/shared/versions"
	v1Publish "github.com/quinntas/visage/api/v1/implementations/publish"
	v1Subscribe "github.com/quinntas/visage/api/v1/implementations/subscribe"
	"github.com/quinntas/visage/internal/api"
)

func Create() api.VersionRouter {
	return api.NewVersionRouter(
		versions.V1,
		api.ImplMap{
			implementations.PUBLISH:   v1Publish.NewPublish(),
			implementations.SUBSCRIBE: v1Subscribe.NewSubscribe(),
		},
	)
}
