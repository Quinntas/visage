package v1Publish

import (
	"context"
	"errors"
	v1Subscribe "github.com/quinntas/visage/api/v1/implementations/subscribe"
	pubsubProtocol "github.com/quinntas/visage/api/v1/protocols/pubsub"
	"github.com/quinntas/visage/internal/api/responses"
	"net"
)

type Publish struct {
	ctx context.Context
}

func NewPublish(ctx context.Context) *Publish {
	return &Publish{
		ctx: ctx,
	}
}

func (p *Publish) Call(content *[]byte, conn net.Conn) []byte {
	channelName, payload, err := pubsubProtocol.ParseContent(content)
	if err != nil {
		return responses.ErrorResponse(err)
	}

	channelNameString := string(channelName)

	channel, ok := p.ctx.Value(v1Subscribe.ACTIVE_CHANNELS_CTX_KEY).(v1Subscribe.ChannelMap)[channelNameString]
	if !ok {
		return responses.ErrorResponse(errors.New("channel not found"))
	}

	channel.SendAll(payload)

	return responses.Ok()
}
