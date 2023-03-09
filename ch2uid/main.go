package main

import (
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	cnt := 1
	n := maelstrom.NewNode()
	n.Handle("init", func(msg maelstrom.Message) error {
		defer func() { cnt++ }()
		var reply maelstrom.MessageBody
		reply.Type = "init_ok"
		reply.MsgID = cnt
		return n.Reply(msg, reply)
	})
	n.Handle("generate", func(msg maelstrom.Message) error {
		defer func() { cnt++ }()
		r := make(map[string]any)
		r["type"] = "generate_ok"
		r["msg_id"] = cnt
		r["id"] = fmt.Sprintf("%s_%d", n.ID(), cnt)
		return n.Reply(msg, r)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
