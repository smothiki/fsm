package main

//example1  assumes no error scenario starts with idle state and ends with close event
// you can see expect error is given false for each event
func example1() {
	fsm := Newsfsm("idle")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", false, 0)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", false, 0)
	eventexecutor(fsm, "close", false, 0)
}

//example2  assumes  error scenario starts with idle state and ends with close event
// you can see expect error is given true for  event makePlan and executeJobs has retries 1
func example2() {
	fsm := Newsfsm("idle")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", true, 1)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", true, 1)
	eventexecutor(fsm, "close", false, 0)
}

//example3  assumes  error scenario starts with idle state and ends with close event
// you can see expect error is given true for  event makePlan retries2 and executeJobs has retries 1
func example3() {
	fsm := Newsfsm("idle")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", true, 2)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", true, 1)
	eventexecutor(fsm, "close", false, 0)
}

//example4  assumes  error scenario starts with  state executeplan. So init and makePlan are NoOp ends with close event
// you can see expect error is given true for  event makePlan retries2 and executeJobs has retries 1
func example4() {
	fsm := Newsfsm("executeplan")
	eventexecutor(fsm, "init", false, 0)
	eventexecutor(fsm, "makePlan", true, 2)
	eventexecutor(fsm, "planexecute", false, 0)
	eventexecutor(fsm, "executeJobs", true, 1)
	eventexecutor(fsm, "close", false, 0)
}
