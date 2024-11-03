package main

import (
	"fmt"
	"root/brute_force"
	s "root/server"
	"runtime"
)

func Exist(c *s.Ctx) {
	format := fmt.Sprintf("cmd: exist\n cors: %d", runtime.NumCPU())
	c.ResWriter(format)
}

func main() {
	//system part
	node := brute_force.NewNode()

	//protocol part
	mux := s.NewRouter()

	mux.HandleFunc("start", node.Start)
	mux.HandleFunc("cancel", node.Cancel)
	mux.HandleFunc("exist", Exist)

	s.ListenAndServe(":5000", mux)
}
