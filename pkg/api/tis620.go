package api

import "unicode/utf8"

// constant variable
const (
	OFFSET = 0xd60
	WIDTH  = 3
)

// ToUTF8 thai industrial standard 620-2533
func ToUTF8(tis620bytes []byte) []byte {
	l := findOutputLength(tis620bytes)
	output := make([]byte, l)

	index := 0
	buffer := make([]byte, WIDTH)
	for _, c := range tis620bytes {
		if !isThaiChar(c) {
			output[index] = c

			index++
		} else {
			utf8.EncodeRune(buffer, int32(c)+OFFSET)
			output[index] = buffer[0]
			output[index+1] = buffer[1]
			output[index+2] = buffer[2]

			index += 3
		}
	}
	return output
}

func findOutputLength(tis620bytes []byte) int {
	outputLen := 0
	for i := range tis620bytes {
		if isThaiChar(tis620bytes[i]) {
			outputLen += WIDTH //always 3 bytes for thai char
		} else {
			outputLen++
		}
	}
	return outputLen
}

func isThaiChar(c byte) bool {
	return (c >= 0xA1 && c <= 0xDA) || (c >= 0xDF && c <= 0xFB)
}
