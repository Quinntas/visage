package protocol

type Impl interface {
	Call(payload []byte) []byte
}

type VersionRouter struct {
	Id    int8
	Impls map[int8]Impl
}

func NewVersionRouter(id int8, implMap map[int8]Impl) VersionRouter {
	return VersionRouter{id, implMap}
}
