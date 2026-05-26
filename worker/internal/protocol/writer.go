package protocol

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

func WriteRCOUNT(conn net.Conn, count map[string]uint64) {
	fmt.Printf("Writing report: %d\n", len(count))

	writterBuf := bufio.NewWriterSize(conn, 64*1024)
	defer writterBuf.Flush()

	var longBuf [2]byte
	var countBuf [8]byte

	for word, wordCount := range count {
		wordBuf := []byte(word)
		if len(wordBuf) > 0xFFFF {
			fmt.Printf("Error: word too long to report: %s\n", word)
        	continue
      	}
		
		binary.BigEndian.PutUint16(longBuf[:], uint16(len(wordBuf)))
		binary.BigEndian.PutUint64(countBuf[:], uint64(wordCount))

		writterBuf.Write(longBuf[:])
		writterBuf.Write(wordBuf)
		writterBuf.Write(countBuf[:])
	}
}