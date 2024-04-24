package v1Subscribe

import (
	"fmt"
	pubsubProtocol "github.com/quinntas/visage/api/v1/protocols/pubsub"
	"github.com/quinntas/visage/internal/api/responses"
)

type Subscribe struct {
}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

func (p *Subscribe) Call(content *[]byte) []byte {
	channel, payload, err := pubsubProtocol.ParseContent(content)
	if err != nil {
		return responses.ErrorResponse(err)
	}

	fmt.Println(string(channel))
	fmt.Println(string(payload))

	return responses.Ok()
}
