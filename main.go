package main

import (
	"root/brute_force"
	s "root/server"
)

func Exist(c *s.Ctx) {
	c.ResWriter("cmd: exist\n")
}

func main() {
	//system part
	bf := brute_force.New()
	node := brute_force.NewNode(bf)

	//protocol part
	mux := s.NewRouter()

	mux.HandleFunc("start", node.Start)
	mux.HandleFunc("cancel", node.Cancel)
	mux.HandleFunc("exist", Exist)

	s.ListenAndServe(":5000", mux)
}
