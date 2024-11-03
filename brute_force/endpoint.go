package brute_force

import (
	"fmt"
	s "root/server"
	"time"
)

func (n *Node) Cancel(c *s.Ctx) {
	if !n.status {
		c.ResWriter("cmd: info\n Msg:\r Server is not active\n")
		return
	}

	n.CancelFn()
	fmt.Println("my cancel") //test info
	c.ResWriter("cmd: info\n Msg:\r Process is cancel!\n")
}

func (n *Node) Start(c *s.Ctx) {
	n.status = true
	start := time.Now()
	// mod := c.GetHeader()["mod"]

	cors := 6
	hash := "JankoKondic"

	pointer := 0

	cancel := NewCancel()
	n.UpdateCancelManager(cancel)

	for range cors {
		go n.Worker(hash, FindCombination(pointer*10_000_000_000), cancel)
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
		}
	}

	format := fmt.Sprintf("cmd: end\n iteration: %d\n found: %s\n duration: %s\n", numberOfIteration, found, time.Since(start).String())
	c.ResWriter(format)
	n.status = false
}
