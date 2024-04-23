package v1Router

import (
	"github.com/quinntas/visage/api/shared/versions"
	v1Publish "github.com/quinntas/visage/api/v1/implementations/publish"
	"github.com/quinntas/visage/internal/api"
)

const (
	PUBLISH = iota
)

func Create() api.VersionRouter {
	return api.NewVersionRouter(
		versions.V1,
		map[uint8]api.Impl{
			PUBLISH: v1Publish.NewPublish(),
		},
	)
}
