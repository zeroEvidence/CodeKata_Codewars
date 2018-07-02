package kata_string_constructing

import (
	"bytes"
)

func StringConstructing(a, s string) int {
	aBytes := []byte(a)
	sBytes := []byte(s)
	opsCount := int(0)
	currentBytes := make([]byte, 0)

	for false == bytes.Equal(currentBytes, sBytes) {

		currentBytes = append(currentBytes, aBytes...)
		opsCount++

		for byteIndex := 0; byteIndex < len(currentBytes); {
			if byteIndex >= len(sBytes) || currentBytes[byteIndex] != sBytes[byteIndex] {
				currentBytes = append(currentBytes[:byteIndex], currentBytes[byteIndex+1:]...)
				opsCount++
			} else {
				byteIndex++
			}
		}
	}

	return opsCount
}
