package main

import (
	"root/brute_force"
	s "root/server"
	"runtime"
)

func Exist(c *s.Ctx) {
	c.ResWriter("cmd: exist\n")
}

func (n *Node) Cancel(c *s.Ctx) {
	n.processController <- struct{}{}
	c.ResWriter("cmd: info\n Msg:\r Process is cancel!\n")
}

func (n *Node) Start(c *s.Ctx) {
	// mod := c.GetHeader()["mod"]
loop:
	for {
		select {
		case <-n.processController:
			break loop
		default:
			for index := 0; index*10_000_000_000 < 1_000_000_000_000; index++ {
				n.Send(brute_force.NewDataStream("JankoKondic", brute_force.FindCombination(index*10_000_000_000)))
			}
		}
	}

	n.work = false
}

type Node struct {
	*brute_force.BruteForce
	processController chan struct{}
	work              bool
}

func NewNode(brutForce *brute_force.BruteForce) Node {
	return Node{
		BruteForce:        brutForce,
		processController: make(chan struct{}),
	}
}

func main() {
	//system part
	bf := brute_force.New()
	node := NewNode(bf)

	for index := range runtime.NumCPU() {
		go bf.Worker(index)
	}

	//protocol part
	mux := s.NewRouter()

	mux.HandleFunc("start", node.Start)
	mux.HandleFunc("cancel", node.Cancel)
	mux.HandleFunc("exist", Exist)

	s.ListenAndServe(":5000", mux)
}
