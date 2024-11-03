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

// func ParseMod(header map[string]string) (int, error) {
// 	mod, ok := header["mod"]
// 	if !ok {
// 		return -1, fmt.Errorf("cant read header [mod]")
// 	}

// 	if mod == "single" {
// 		return 1, nil
// 	}

// 	return runtime.NumCPU(), nil
// }

// func ParseStartPosition(header map[string]string) (int, error) {
// 	number, ok := header["point"]
// 	if !ok {
// 		return -1, fmt.Errorf("cant read header [mod]")
// 	}

// 	startPointer, err := strconv.Atoi(number)
// 	if err != nil {
// 		return -1, fmt.Errorf(err.Error())
// 	}

// 	return startPointer, nil
// }

// func StartParser(header map[string]string) (int, int, string, error) {
// 	mod, err := ParseMod(header)
// 	if err != nil {
// 		return 0, 0, "", err
// 	}

// 	hash, ok := header["hash"]
// 	if !ok {
// 		return 0, 0, "", err
// 	}

// 	startPoint, err := ParseStartPosition(header)
// 	if err != nil {
// 		return 0, 0, "", err
// 	}

// 	return mod, startPoint, hash, nil
// }

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
		go n.Worker(payload.Hash, FindCombination(payload.Pointer*10_000_000_000), cancel)
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

	format := fmt.Sprintf("cmd: end\n iteration: %d\n found: %s\n", numberOfIteration, found)
	c.ResWriter(format)
	n.status = false
}
