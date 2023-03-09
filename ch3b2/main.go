package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastReq struct {
	maelstrom.MessageBody
	Message int `json:"message"`
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
	n         *maelstrom.Node
	storage   map[int]struct{} // because "The value is always an integer and it is unique for each message from Maelstrom."
	storageRW sync.RWMutex
}

func (s *server) appendValue(v int) (alreadyExist bool) {
	s.storageRW.Lock()
	_, ok := s.storage[v]
	if ok {
		s.storageRW.Unlock()
		return true
	}
	s.storage[v] = struct{}{}
	s.storageRW.Unlock()
	return false
}

func (s *server) getValues() []int { // order does not matter
	s.storageRW.RLock()
	ret := make([]int, len(s.storage))
	idx := 0
	for v := range s.storage {
		ret[idx] = v
		idx++
	}
	s.storageRW.RUnlock()
	return ret
}

func (s *server) Broadcast(msg maelstrom.Message) error {
	var req BroadcastReq
	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		return fmt.Errorf("unmarshal req err: %+v", err)
	}

	alreadyExist := s.appendValue(req.Message)
	if !alreadyExist { // only broadcast new values
		s.doBroadcast(msg.Src, req)
	}

	rsp := maelstrom.MessageBody{Type: "broadcast_ok"}
	return s.n.Reply(msg, rsp)
}

func (s *server) doBroadcast(from string, req any) {
	// just ignore topology
	for _, node := range s.n.NodeIDs() {
		if node == s.n.ID() { // do not send to self
			continue
		}
		if node == from { // do not send back
			continue
		}
		s.n.Send(node, req)
	}
}

func (s *server) Read(msg maelstrom.Message) error {
	rsp := ReadRsp{
		MessageBody: maelstrom.MessageBody{
			Type: "read_ok",
		},
		Messages: s.getValues(),
	}
	return s.n.Reply(msg, rsp)
}

func (s *server) Topology(msg maelstrom.Message) error {
	rsp := maelstrom.MessageBody{Type: "topology_ok"}
	return s.n.Reply(msg, rsp)
}

func main() {
	s := &server{n: maelstrom.NewNode(), storage: make(map[int]struct{})}
	s.n.Handle("broadcast", s.Broadcast)
	s.n.Handle("read", s.Read)
	s.n.Handle("topology", s.Topology)

	if err := s.n.Run(); err != nil {
		log.Fatal(err)
	}
}
