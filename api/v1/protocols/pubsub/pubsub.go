package pubsubProtocol

import (
	"errors"
	"github.com/quinntas/visage/utils/byteUtils"
)

func ParseContentHeader(contentHeader []byte) (uint16, uint16, error) {
	contentHeaderLen := len(contentHeader)

	if contentHeaderLen != 4 {
		return 0, 0, errors.New("not enough bytes")
	}

	contentHeaderChannelLength := uint16(contentHeader[0])<<8 | uint16(contentHeader[1])

	contentHeaderPayloadLength := uint16(contentHeader[2])<<8 | uint16(contentHeader[3])

	return contentHeaderChannelLength, contentHeaderPayloadLength, nil
}

func ParseContent(content *[]byte) ([]byte, []byte, error) {
	contentHeader, err := byteUtils.ReadByteArr(0, CONTENT_HEADER_LENGTH, content)
	if err != nil {
		return nil, nil, err
	}

	contentHeaderChannelLength, contentHeaderPayloadLength, err := ParseContentHeader(contentHeader)
	if err != nil {
		return nil, nil, err
	}

	startChannelIndex := CONTENT_HEADER_LENGTH + int(contentHeaderChannelLength)
	channel, err := byteUtils.ReadByteArr(0, startChannelIndex, content)
	if err != nil {
		return nil, nil, err
	}

	endPayloadIndex := int(contentHeaderPayloadLength) + startChannelIndex
	payload, err := byteUtils.ReadByteArr(startChannelIndex, endPayloadIndex, content)
	if err != nil {
		return nil, nil, err
	}

	return channel, payload, nil
}
