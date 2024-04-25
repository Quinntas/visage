package v1Subscribe

import (
	"context"
	pubsubProtocol "github.com/quinntas/visage/api/v1/protocols/pubsub"
	"github.com/quinntas/visage/internal/api/responses"
	"net"
)

type Subscribe struct {
	ctx context.Context
}

func NewSubscribe(ctx context.Context) *Subscribe {
	return &Subscribe{
		ctx: ctx,
	}
}

func ApplySubscribeContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ACTIVE_CHANNELS_CTX_KEY, ChannelMap{})
}

// TODO: race conditions
// TODO: make the sub value a uuidv4 so if the client disconnect it may reconnect and receive all previus messages
func (p *Subscribe) Call(content *[]byte, conn net.Conn) []byte {
	channelName, _, err := pubsubProtocol.ParseContent(content)
	if err != nil {
		return responses.ErrorResponse(err)
	}

	channelNameString := string(channelName)

	channel, ok := p.ctx.Value(ACTIVE_CHANNELS_CTX_KEY).(ChannelMap)[channelNameString]
	if !ok {
		channel = NewChannel(channelNameString)
	}

	err = channel.AppendConn(conn)
	if err != nil {
		return responses.ErrorResponse(err)
	}

	// TODO: maybe change this ? minor perf improv
	p.ctx.Value(ACTIVE_CHANNELS_CTX_KEY).(ChannelMap)[channelNameString] = channel

	// TODO: dont like this
	err = channel.HoldConnectionRoutine(conn)
	if err != nil {
		return responses.ErrorResponse(err)
	}

	return responses.Ok()
}
