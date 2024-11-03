package brute_force

type Node struct {
	status bool
	*BruteForce
	*CancelManager
}

func (n *Node) UpdateCancelManager(c *CancelManager) {
	n.CancelManager = c
}

func NewNode() Node {
	return Node{
		BruteForce: New(),
	}
}
