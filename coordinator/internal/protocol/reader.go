package protocol

import (
	"bufio"
	"encoding/binary"
	"io"
)

func ReadReportChunkId(reader *bufio.Reader) (uint32, error) {
	var chunkIdBuf [4]byte

	if _, err := io.ReadFull(reader, chunkIdBuf[:]); err != nil {
		return 0, err
	}

	chunkId := binary.BigEndian.Uint32(chunkIdBuf[:])

	return chunkId, nil
}

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