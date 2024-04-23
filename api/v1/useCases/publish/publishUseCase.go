package v1PublishUseCase

type PublishUseCase struct {
}

func NewPublishUseCase() *PublishUseCase {
	return &PublishUseCase{}
}

func (p *PublishUseCase) Call(payload []byte) []byte {
	return payload
}
