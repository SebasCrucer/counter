package protocol

import (
	"fmt"
	"net"
)

func WriteRCOUNT(conn net.Conn, count int) {
	conn.Write([]byte(" "))
	m := fmt.Sprintf(RCOUNT, count)
	conn.Write([]byte(m))
}