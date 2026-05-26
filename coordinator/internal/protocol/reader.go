package protocol

import (
	"bufio"
	"encoding/binary"
	"io"
)

func ReadReport(reader *bufio.Reader) (string, uint64, error) {
	var longBuf [2]byte

	if _, err := io.ReadFull(reader, longBuf[:]); err != nil {
		return "", 0, err
	}

	long := binary.BigEndian.Uint16((longBuf[:]))

	wordBytes := make([]byte, long)

	if _, err := io.ReadFull(reader, wordBytes); err != nil {
		return "", 0, err
	}

	word := string(wordBytes)

	var countBytes [8]byte

	if _, err := io.ReadFull(reader, countBytes[:]); err != nil {
		return "", 0, err
	}

	count := binary.BigEndian.Uint64(countBytes[:])

	return word, count, nil
}