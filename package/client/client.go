package client

import (
	"errors"
	"fmt"
	"net"
	"time"
)

var (
	ServerHost = "localhost"
	ServerPort = "2222"

	ReadTimeoutMs = 30000 // in ms
)

const ServerType = "udp"

var (
	ErrFailedToSetReadDeadline = errors.New("failed to set conn read deadline")
)

func MustInit() UdpConn {
	serverStr := fmt.Sprintf("%s:%s", ServerHost, ServerPort)
	udpServer, err := net.ResolveUDPAddr(ServerType, serverStr)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Established UDP connection to %s\n", serverStr)
	return UdpConn{conn}
}

type UdpConn struct {
	x net.Conn
}

// Apply read deadline for timeouts
func (u *UdpConn) Write(b []byte) (n int, err error) {
	return u.x.Write(b)
}

func (u *UdpConn) Read(b []byte) (int, error) {
	if err := u.x.SetReadDeadline(time.Now().Add(time.Duration(ReadTimeoutMs) * time.Millisecond)); err != nil {
		return 0, ErrFailedToSetReadDeadline
	}
	return u.x.Read(b)
}

func (u *UdpConn) Close() error {
	return u.x.Close()
}
