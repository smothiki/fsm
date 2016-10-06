package main

import "github.com/looplab/fsm"

type Job struct {
	jobspc string //generally struct
	FSM    *fsm.FSM
}

// go run examples/genesis/simple.go examples/genesis/util.go examples/genesis/examples.go
func main() {
	example4()
}
