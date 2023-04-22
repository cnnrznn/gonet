package gonet

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBigSend(t *testing.T) {
	client := New(":9991")
	server := New(":9998")
	sync := make(chan error)
	size := 16000

	msg := make([]byte, size)
	n, err := rand.Read(msg)
	if err != nil {
		t.Error(err)
	}
	if n != size {
		t.Error(fmt.Errorf("wrong payload size"))
	}

	go func() {
		_, err := server.Recv(size)
		sync <- err
	}()

	err = client.Send(msg, "localhost:9998")
	if err != nil {
		t.Error(err)
	}

	for err := range sync {
		if err != nil {
			t.Error(err)
		}
	}
}
