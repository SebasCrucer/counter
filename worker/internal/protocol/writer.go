package protocol

import (
	"encoding/binary"
	"fmt"
	"net"
)

func WriteRCOUNT(conn net.Conn, count map[string]uint64) {
	fmt.Printf("Writing report: %d\n", len(count))
	for word, wordCount := range count {
		wordBytes := []byte(word)
		if len(wordBytes) > 0xFFFF {
			fmt.Printf("Error: word too long to report: %s\n", word)
        	return
      	}
		
		buf := make([]byte, 2+len(wordBytes)+8)
		
		binary.BigEndian.PutUint16(buf[0:2], uint16(len(wordBytes)))
		copy(buf[2:2+len(wordBytes)], wordBytes)
		binary.BigEndian.PutUint64(buf[2+len(wordBytes):], uint64(wordCount))

		if _, err := conn.Write(buf); err != nil {
			fmt.Printf("Error writing report: %s\n", err)
			return
		}
	}
}