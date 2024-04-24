package byteUtils

import "errors"

func ReadByteArr(startIndex int, endIndex int, bytes *[]byte) ([]byte, error) {
	bytesLen := len(*bytes)

	if bytesLen < startIndex {
		return nil, errors.New("bytes length is less than startIndex")
	} else if bytesLen < endIndex {
		return nil, errors.New("bytes length is less than endIndex")
	}

	return (*bytes)[startIndex:endIndex], nil
}
