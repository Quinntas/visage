package reponses

func ByteResponse(payload []byte) []byte {
	return payload
}

func Ok() []byte {
	return ByteResponse([]byte("Ok"))
}
