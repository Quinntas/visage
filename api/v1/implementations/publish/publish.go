package v1Publish

import (
	"fmt"
	pubsubProtocol "github.com/quinntas/visage/api/v1/protocols/pubsub"
	"github.com/quinntas/visage/internal/api/responses"
)

type Publish struct {
}

func NewPublish() *Publish {
	return &Publish{}
}

func (p *Publish) Call(content *[]byte) []byte {
	channel, payload, err := pubsubProtocol.ParseContent(content)
	if err != nil {
		return responses.ErrorResponse(err)
	}

	fmt.Println(channel)
	fmt.Println(payload)

	return responses.Ok()
}
