package protocol

import (
	"net"
)

func WriteGENDCODE(conn net.Conn) {
	conn.Write([]byte(" "))
	conn.Write([]byte(GENDCODE))
}