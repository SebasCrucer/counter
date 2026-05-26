package protocol

import (
	"net"
)

func WriteWGOCODE(conn net.Conn) {
	conn.Write([]byte{byte(WGOCODE)})
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