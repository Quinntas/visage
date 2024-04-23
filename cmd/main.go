package main

import (
	"fmt"
	sharedRouter "github.com/quinntas/visage/api/shared/infra/router"
	"net"
	"os"
)

/*
protocol
a b c d
a - versioning - 1 byte
b - command - 1 byte
c - length - 2 bytes
d - content
*/

type RawProtocol struct {
	version uint8
	command uint8
	length  int16
	content []byte
}

func NewRawProtocol(
	version uint8,
	command uint8,
	length int16,
	content []byte,
) RawProtocol {
	return RawProtocol{
		version: version,
		command: command,
		length:  length,
		content: content,
	}
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	header := make([]byte, 4)
	_, err := conn.Read(header)
	if err != nil {
		fmt.Println("Error reading header:", err.Error())
		return
	}

	length := int16(header[2])<<8 | int16(header[3])

	content := make([]byte, length)
	totalRead := 0
	for totalRead < int(length) {
		n, err := conn.Read(content[totalRead:])
		if err != nil {
			fmt.Println("Error reading content:", err.Error())
			return
		}
		totalRead += n
	}

	message := NewRawProtocol(
		header[0],
		header[1],
		length,
		content,
	)

	a := sharedRouter.Create()

	fmt.Println(string(a[0].Impls[0].Call(message.content)))

	fmt.Println("Version:", message.version)
	fmt.Println("Command:", message.command)
	fmt.Println("Length:", message.length)
	fmt.Println("Content:", string(message.content))

	_, err = conn.Write([]byte("OK"))
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}
}
