package v1Publish

import (
	"fmt"
	"github.com/quinntas/visage/internal/api/reponses"
)

type Publish struct {
}

func NewPublish() *Publish {
	return &Publish{}
}

func (p *Publish) Call(payload []byte) []byte {
	fmt.Println(string(payload))

	return reponses.Ok()
}
