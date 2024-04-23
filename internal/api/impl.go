package api

type Impl interface {
	Call(payload []byte) []byte
}
