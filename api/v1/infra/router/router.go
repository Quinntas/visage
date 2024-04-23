package v1Router

import (
	"github.com/quinntas/visage/api/shared/sharedVersions"
	v1PublishUseCase "github.com/quinntas/visage/api/v1/useCases/publish"
	"github.com/quinntas/visage/internal/protocol"
)

const (
	PUBLISH = iota
)

func Create() protocol.VersionRouter {
	return protocol.NewVersionRouter(
		sharedVersions.V1,
		map[int8]protocol.Impl{
			PUBLISH: v1PublishUseCase.NewPublishUseCase(),
		},
	)
}
