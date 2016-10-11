package main

import (
	"flag"

	"github.com/looplab/fsm"
)

type Job struct {
	jobspc string //generally struct
	FSM    *fsm.FSM
}

// go run examples/genesis/simple.go examples/genesis/util.go examples/genesis/examples.go
func main() {
	example := flag.Int("example", 1, "a string")
	flag.Parse()
	switch *example {
	case 1:
		example1()
		break
	case 2:
		example2()
		break
	case 3:
		example3()
		break
	case 4:
		example4()
	}
}
