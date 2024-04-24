package api

type ImplMap map[uint8]Impl

type RouterMap map[uint8]VersionRouter

const (
	ROUTERS_CTX_KEY = "routers"
)

type VersionRouter struct {
	Id    uint8
	Impls ImplMap
}

func NewVersionRouter(id uint8, implMap ImplMap) VersionRouter {
	return VersionRouter{id, implMap}
}
