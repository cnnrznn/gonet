package gonet

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBigSend(t *testing.T) {
	client := New("localhost:9971", 3)
	server := New("localhost:9972", 3)
	sync := make(chan error)
	size := 8000000000

	go func() {
		_, err := server.Recv(size)
		sync <- err
		close(sync)
	}()

	msg := make([]byte, size)
	n, err := rand.Read(msg)
	if err != nil {
		t.Error(err)
	}
	if n != size {
		t.Error(fmt.Errorf("wrong payload size"))
	}

	err = client.Send(msg, "localhost:9972")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for err := range sync {
		if err != nil {
			t.Error(err)
		}
	}
}
