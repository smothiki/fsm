package main

func example1() {
	fsm := Newsfsm("idle")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", false, 0)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", false, 0)
	eventexecutor(fsm, "close", false, 0)
}

func example2() {
	fsm := Newsfsm("idle")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", true, 1)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", true, 1)
	eventexecutor(fsm, "close", false, 0)
}

func example3() {
	fsm := Newsfsm("idle")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", true, 2)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", true, 1)
	eventexecutor(fsm, "close", false, 0)
}

func example4() {
	fsm := Newsfsm("executeplan")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", true, 2)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", true, 1)
	eventexecutor(fsm, "close", false, 0)
}
