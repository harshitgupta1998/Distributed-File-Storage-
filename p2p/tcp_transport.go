package p2p

import (
	"net"
	"sync"
	"fmt"
)

type TCPTransport struct {
	listenAddress string
	listener   net.Listener
	mu 		   sync.RWMutex
	peers	  map[net.Addr]Peer
}

func NewTCPTransportOptions(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAccepting()
	return nil
}


func (t *TCPTransport) startAccepting() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}
		go t.handleConnection(conn)
	}
}
func (t *TCPTransport) handleConnection(conn net.Conn) {
	fmt.Println("Accepted connection from", conn)
}