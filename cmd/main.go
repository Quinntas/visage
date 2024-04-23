package main

import (
	"context"
	"flag"
	"fmt"
	sharedRouter "github.com/quinntas/visage/api/shared/infra/router"
	"github.com/quinntas/visage/internal/api"
	"net"
	"os"
)

func main() {
	hostPtr := flag.String("host", "localhost", "tcp host")
	portPtr := flag.String("port", "6969", "tcp port")

	args := api.NewArgs(*hostPtr, *portPtr)

	l, err := net.Listen("tcp", args.Addr())
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	ctx := sharedRouter.CreateWithContext(context.Background())

	fmt.Println("Listening on " + args.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go api.HandleConnectedClient(conn, ctx)
	}
}
