package simulator

import (
	"fmt"
	"testing"
)

func TestSimulator2Nodes(t *testing.T) {
	node := New([]string{
		"localhost:2",
	})

	msg := []byte("Howdy partner!")
	size := len(msg)

	node.Send(msg, "localhost:2")
	bs, _ := node.Recv("localhost:2", size)

	equal := true
	for i, b := range msg {
		if b != bs[i] {
			equal = false
		}
	}

	if !equal {
		t.Error(fmt.Errorf("message received different from expected: %v | %v", string(msg), string(bs)))
	}
}
