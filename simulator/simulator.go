package simulator

import "fmt"

type Message struct {
	Data []byte
}

type SimNode struct {
	pipes map[string]chan Message
}

func New(addrs []string) *SimNode {
	sim := &SimNode{
		pipes: make(map[string]chan Message),
	}

	for _, addr := range addrs {
		sim.pipes[addr] = make(chan Message, 4096)
	}

	return sim
}

func (n *SimNode) Send(data []byte, addr string) error {
	n.pipes[addr] <- Message{
		Data: data,
	}
	return nil
}

func (n *SimNode) Recv(addr string, size int) ([]byte, error) {
	msg := <-n.pipes[addr]

	if len(msg.Data) != size {
		return nil, fmt.Errorf("incorrect number of bytes received")
	}

	return msg.Data, nil
}
