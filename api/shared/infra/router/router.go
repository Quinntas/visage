package sharedRouter

import (
	"github.com/quinntas/visage/api/shared/sharedVersions"
	v1Router "github.com/quinntas/visage/api/v1/infra/router"
	"github.com/quinntas/visage/internal/protocol"
)

func Create() map[int8]protocol.VersionRouter {
	return map[int8]protocol.VersionRouter{
		sharedVersions.V1: v1Router.Create(),
	}
}
