package client

import (
	"net"
	"testing"
	"time"
)

// createTestUDPServer creates a simple instance of a UDP server that echos back sent packets
func createTestUDPServer() {
	udpServer, err := net.ListenPacket("udp", "localhost:2222")
	if err != nil {
		panic(err)
	}
	defer udpServer.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		if _, err := udpServer.WriteTo([]byte("Server Echo: "+string(buf[:n])), addr); err != nil {
			panic(err)
		}
	}
}

func TestMustInit(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr error
	}{
		{
			name:    "read write simple",
			msg:     "apple boy cat 123",
			wantErr: nil,
		},
	}
	go createTestUDPServer()
	// wait for server to start
	time.Sleep(1 * time.Second)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := MustInit()
			if _, err := conn.Write([]byte(tt.msg)); err != tt.wantErr {
				t.Fatalf("expected err: %v, got err: %v", tt.wantErr, err)
			}

			buffer := make([]byte, 1024)
			mLen, err := conn.Read(buffer, time.Now().Add(time.Minute))
			if err != nil {
				t.Fatal(err)
			}
			gotStr := string(buffer[:mLen])
			expStr := "Server Echo: " + tt.msg
			if gotStr != expStr {
				t.Fatalf("expected string: %s, got %s", expStr, gotStr)
			}
			conn.Close()
		})
	}
}
