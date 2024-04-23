package api

type Protocol struct {
	Version uint8
	Command uint8
	Length  uint16
	Content []byte
}

func NewProtocol(
	version uint8,
	command uint8,
	length uint16,
	content []byte,
) Protocol {
	return Protocol{
		Version: version,
		Command: command,
		Length:  length,
		Content: content,
	}
}
