package brute_force

import (
	"encoding/json"
	"fmt"
	s "root/server"
	"runtime"
)

func (n *Node) Cancel(c *s.Ctx) {
	if !n.status {
		c.ResWriter("cmd: info\n Msg:\r Server is not active\n")
		return
	}

	n.CancelFn()
	fmt.Println("my cancel") //Delete later
	c.ResWriter("cmd: info\n Msg:\r Process is cancel!\n")
}

func (n *Node) Start(c *s.Ctx) {
	payload := &struct {
		Pointer int    `json:"pointer"`
		Hash    string `json:"hash"`
		Mod     string `json:"mod"`
	}{}

	if err := json.Unmarshal([]byte(c.GetHeader()["body"]), payload); err != nil {
		c.ResWriter(err.Error())
		return
	}

	var cors int = runtime.NumCPU()
	if payload.Mod == "single" {
		cors = 1
	}

	n.status = true
	fmt.Println(payload.Pointer, payload.Hash) //Delete later

	cancel := NewCancel()
	n.UpdateCancelManager(cancel)

	for range cors {
		go n.Worker(payload.Hash, FindCombination(payload.Pointer), cancel)
		payload.Pointer++
	}

	var (
		numberOfIteration int
		found             string
	)

	for range cors {
		r := n.GetResponseCh()
		numberOfIteration += r.GetIteration()

		if r.GetPassword() == payload.Hash {
			found = payload.Hash
			fmt.Println("Found") //Delete later
		}
	}

	c.ResWriter(StartFormat(cancel.GetState(), numberOfIteration, found))
	n.status = false
}

func StartFormat(status bool, numberOfIteration int, found string) string {
	if status {
		return fmt.Sprintf("cmd: end\n iteration: %d\n found: %s\n", numberOfIteration, found)
	}

	return fmt.Sprintf("cmd: cancel\n")
}
