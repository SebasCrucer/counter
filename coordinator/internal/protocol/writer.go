package protocol

import (
	"encoding/binary"
	"net"
)

func WriteWGOCODE(conn net.Conn, chunkIndex uint32) {
	var buf [5]byte
	buf[0] = byte(WGOCODE)
	binary.BigEndian.PutUint32(buf[1:], chunkIndex)
	conn.Write(buf[:])
}

func WriteWGENDCODE(conn net.Conn) {
	conn.Write([]byte{byte(WGENDCODE)})
}

func WriteROK(conn net.Conn) {
	conn.Write([]byte{byte(ROK)})
}

func WriteRERROR(conn net.Conn) {
	conn.Write([]byte{byte(RERROR)})
}