package v1Subscribe

import (
	"errors"
	"net"
)

type ChannelMap map[string]Channel

// TODO: implement queues to each connetion, so it can receive all data it didnt receive in the reconection
type Channel struct {
	Name        string
	ActiveConns []net.Conn
}

func (c *Channel) SendAll(bytes []byte) {
	for _, conn := range c.ActiveConns {
		_, err := conn.Write(bytes)
		if err != nil {
			go c.RemoveConn(conn)
			continue
		}
	}
}

// TODO: race conditions
func (c *Channel) RemoveConn(conn net.Conn) {
	for _, ActiveConns := range c.ActiveConns {
		if ActiveConns.RemoteAddr() == conn.RemoteAddr() {
			c.ActiveConns = append(c.ActiveConns[:0], c.ActiveConns[1:]...)
		}
	}
}

// TODO: refactor this
// TODO: add healthCheck timeout
func (c *Channel) HoldConnectionRoutine(conn net.Conn) error {
	for {
		check := make([]byte, len(HEALTHCHECK_STRING))
		totalRead := 0
		for totalRead < len(HEALTHCHECK_STRING) {
			n, err := conn.Read(check[totalRead:])
			if err != nil {
				c.RemoveConn(conn)
				return errors.New("error reading from conn in channel")
			}
			totalRead += n
		}
		if string(check) != HEALTHCHECK_STRING {
			c.RemoveConn(conn)
			return errors.New("health check failed")
		}
	}
}

func (c *Channel) AppendConn(conn net.Conn) error {
	for _, ActiveConns := range c.ActiveConns {
		if ActiveConns.RemoteAddr() == conn.RemoteAddr() {
			return errors.New("already connected")
		}
	}
	c.ActiveConns = append(c.ActiveConns, conn)
	return nil
}

func NewChannel(name string) Channel {
	return Channel{
		Name:        name,
		ActiveConns: []net.Conn{},
	}
}
