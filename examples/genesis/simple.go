package main

import "github.com/looplab/fsm"

type Job struct {
	jobspc string //generally struct
	FSM    *fsm.FSM
}

func main() {
	example3()
}
