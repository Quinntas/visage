package api

import (
	"context"
	"errors"
	"fmt"
	"net"
)

// TODO: buf overflow and shit + module for parsing byte arr
// TODO: receive a result instead of []bytes to not close a connection when .Call() is called :)
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

	ver, ok := ctx.Value(ROUTERS_CTX_KEY).(RouterMap)[protocol.Version]
	if !ok {
		return errors.New("version not found")
	}

	com, ok := ver.Impls[protocol.Command]
	if !ok {
		return errors.New("command not found")
	}

	res := com.Call(&protocol.Content, conn)
	_, err := conn.Write(res)
	if err != nil {
		return err
	}

	return nil
}

// TODO: shitty code
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
