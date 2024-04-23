package api

const (
	ROUTERS_CTX_KEY = "routers"
)

type VersionRouter struct {
	Id    uint8
	Impls map[uint8]Impl
}

func NewVersionRouter(id uint8, implMap map[uint8]Impl) VersionRouter {
	return VersionRouter{id, implMap}
}
