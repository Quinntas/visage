package api

import (
	"context"
	"fmt"
	"net"
)

func handleRequest(header []byte, conn net.Conn, ctx context.Context) error {
	length := uint16(header[2])<<8 | uint16(header[3])

	content := make([]byte, length)
	totalRead := 0
	for totalRead < int(length) {
		n, err := conn.Read(content[totalRead:])
		if err != nil {
			return err
		}
		totalRead += n
	}

	protocol := NewProtocol(
		header[0],
		header[1],
		length,
		content,
	)

	res := ctx.Value(ROUTERS_CTX_KEY).(map[uint8]VersionRouter)[protocol.Version].
		Impls[protocol.Command].
		Call(content)

	_, err := conn.Write(res)
	if err != nil {
		return err
	}

	return nil
}

func HandleConnectedClient(conn net.Conn, ctx context.Context) {
	defer conn.Close()

	header := make([]byte, 4)
	_, err := conn.Read(header)
	if err != nil {
		fmt.Println("Error reading header:", err.Error())
		return
	}

	err = handleRequest(header, conn, ctx)
	if err != nil {
		fmt.Println("Error handling request:", err.Error())
		return
	}
}
