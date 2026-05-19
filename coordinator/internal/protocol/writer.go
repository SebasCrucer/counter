package protocol

import (
	"net"
)

func WriteGENDCODE(conn net.Conn) {
	conn.Write([]byte(" "))
	conn.Write([]byte(GENDCODE))
}

func WriteRERROR(conn net.Conn) {
	conn.Write([]byte(RERROR))
}

func WriteROK(conn net.Conn) {
	conn.Write([]byte(ROK))
}