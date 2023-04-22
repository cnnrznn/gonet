package gonet

import (
	"fmt"
	"net"
)

type Node struct {
	Addr string
}

func New(addr string) *Node {
	return &Node{
		Addr: addr,
	}
}

func (n *Node) Send(data []byte, addr string) error {
	conn, err := net.Dial("tcp", addr)
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

	n, err := conn.Read(data)
	if err != nil {
		return nil, err
	}

	if n != size {
		return nil, fmt.Errorf("bytes expected: %v, read: %v", size, n)
	}

	return data, nil
}
