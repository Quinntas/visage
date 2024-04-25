package api

import "net"

type Impl interface {
	Call(payload *[]byte, conn net.Conn) []byte
}
