package main

import (
	"encoding/json"
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReq struct {
	maelstrom.MessageBody
	Message int  `json:"message"`
	Stop    bool `json:"stop"`
}

type ReadRsp struct {
	maelstrom.MessageBody
	Messages []int `json:"messages"`
}

type TopologyReq struct {
	maelstrom.MessageBody
	Topology map[string][]string `json:"topology"`
}

type server struct {
	n       *maelstrom.Node
	storage []int
}

func (s *server) Broadcast(msg maelstrom.Message) error {
	var req BroadcastReq
	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		return fmt.Errorf("unmarshal req err: %+v", err)
	}

	s.storage = append(s.storage, req.Message)

	if !req.Stop {
		req.Stop = true
		for _, node := range s.n.NodeIDs() {
			if node == s.n.ID() {
				continue
			}
			s.n.Send(node, req)
		}
	}

	rsp := maelstrom.MessageBody{Type: "broadcast_ok"}
	return s.n.Reply(msg, rsp)
}

func (s *server) Read(msg maelstrom.Message) error {
	rsp := ReadRsp{
		MessageBody: maelstrom.MessageBody{
			Type: "read_ok",
		},
		Messages: s.storage,
	}
	return s.n.Reply(msg, rsp)
}

func (s *server) Topology(msg maelstrom.Message) error {
	rsp := maelstrom.MessageBody{Type: "topology_ok"}
	return s.n.Reply(msg, rsp)
}

func main() {
	s := &server{n: maelstrom.NewNode()}
	s.n.Handle("broadcast", s.Broadcast)
	s.n.Handle("read", s.Read)
	s.n.Handle("topology", s.Topology)

	if err := s.n.Run(); err != nil {
		log.Fatal(err)
	}
}
