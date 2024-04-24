package v1Subscribe

import (
	"fmt"
	"github.com/quinntas/visage/internal/api/responses"
)

type Subscribe struct {
}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

// TODO: use the module to parse byte arr
func (p *Subscribe) Call(content []byte) []byte {
	contentChannelLength := uint16(content[0])<<8 | uint16(content[1])
	contentPayloadLength := uint16(content[2])<<8 | uint16(content[3])

	startIndex := 4 + contentChannelLength

	fmt.Println(contentChannelLength, string(content[:startIndex]))
	fmt.Println(contentPayloadLength, string(content[startIndex:contentPayloadLength+startIndex]))

	return responses.Ok()
}
