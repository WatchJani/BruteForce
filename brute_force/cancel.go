package brute_force

import "sync"

type CancelManager struct {
	state bool
	sync.RWMutex
}

func (c *CancelManager) GetState() bool {
	return c.state
}

func NewCancel() *CancelManager {
	return &CancelManager{
		state: true,
	}
}

func (c *CancelManager) CancelFn() {
	c.Lock()
	defer c.Unlock()

	c.state = !c.state
}
