package gonet

import (
	"fmt"
	"net"
	"time"
)

type Node struct {
	Addr        string
	ConnRetries int // number of 'Dial()' retries
}

func New(addr string, retries int) *Node {
	return &Node{
		Addr:        addr,
		ConnRetries: retries,
	}
}

func (n *Node) Send(data []byte, addr string) error {
	var conn net.Conn
	var err error
	for i := 0; i < n.ConnRetries; i++ {
		conn, err = net.Dial("tcp", addr)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return err
	}
	defer conn.Close()

	written := 0
	for written < len(data) {
		n, err := conn.Write(data[written:])
		if err != nil {
			return err
		}

		written += n
	}

	return nil
}

func (node *Node) Recv(size int) ([]byte, error) {
	data := make([]byte, size)

	l, err := net.Listen("tcp", node.Addr)
	if err != nil {
		return nil, err
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	read := 0
	for read < size {
		n, err := conn.Read(data)
		if err != nil {
			return nil, err
		}
		read += n
	}

	if read != size {
		return nil, fmt.Errorf("bytes expected: %v, read: %v", size, read)
	}

	return data, nil
}
