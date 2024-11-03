package main

import (
	"context"
	"fmt"
	"root/brute_force"
	s "root/server"
)

func Exist(c *s.Ctx) {
	c.ResWriter("cmd: exist\n")
}

func (n *Node) Cancel(c *s.Ctx) {
	if n.status {
		c.ResWriter("cmd: info\n Msg:\r Server is not active\n")
	}

	n.cancel()
	fmt.Println("my cancel") //test info
	c.ResWriter("cmd: info\n Msg:\r Process is cancel!\n")
}

func (n *Node) Start(c *s.Ctx) {
	n.status = true
	// mod := c.GetHeader()["mod"]

	cors := 6
	hash := "Janko"

	pointer := 0

	for range cors {
		//password, starting combination, cancel context
		go n.Worker(hash, brute_force.FindCombination(pointer*10_000_000_000), n.CreateContext())
		pointer++
	}

	var (
		numberOfIteration int
		found             string
	)

	for range cors {
		r := n.GetResponseCh()
		numberOfIteration += r.GetIteration()

		if r.GetPassword() == hash {
			found = hash

			fmt.Println("Found")
			n.cancel()
		}
	}

	format := fmt.Sprintf("cmd: end\n iteration: %d\n found: %s", numberOfIteration, found)
	c.ResWriter(format)
}

type Node struct {
	status bool
	*brute_force.BruteForce
	cancel context.CancelFunc
}

func (n *Node) CreateContext() context.Context {
	context, cancel := context.WithCancel(context.Background())
	n.cancel = cancel

	return context
}

func NewNode(brutForce *brute_force.BruteForce) Node {
	return Node{
		BruteForce: brutForce,
	}
}

func main() {
	//system part
	bf := brute_force.New()
	node := NewNode(bf)

	//protocol part
	mux := s.NewRouter()

	mux.HandleFunc("start", node.Start)
	mux.HandleFunc("cancel", node.Cancel)
	mux.HandleFunc("exist", Exist)

	s.ListenAndServe(":5000", mux)
}
