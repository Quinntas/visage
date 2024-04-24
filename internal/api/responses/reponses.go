package responses

// TODO: make some error responses + response codes for better error eval

func ByteResponse(payload []byte) []byte {
	return payload
}

func Ok() []byte {
	return ByteResponse([]byte("Ok"))
}

func ErrorResponse(err error) []byte {
	return ByteResponse([]byte(err.Error()))
}
