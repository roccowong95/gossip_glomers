package main

import (
	"fmt"
	"log"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// UIDMessage wraps id around common msg
type UIDMessage struct {
	maelstrom.MessageBody
	ID any `json:"id"`
}

// UIDServer interface
type UIDServer interface {
	Generate(msg maelstrom.Message) error
}

type naiveServer struct {
	cnt  int64
	node *maelstrom.Node
}

// Generate generate uid
func (s *naiveServer) Generate(msg maelstrom.Message) error {
	defer atomic.AddInt64(&s.cnt, 1)
	var rsp UIDMessage
	rsp.Type = "generate_ok"
	rsp.MsgID = int(s.cnt)
	rsp.ID = fmt.Sprintf("%s_%d", s.node.ID(), s.cnt)
	return s.node.Reply(msg, rsp)
}

func main() {
	n := maelstrom.NewNode()
	// init message is handled by node itself
	s := &naiveServer{node: n}
	n.Handle("generate", s.Generate)
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
