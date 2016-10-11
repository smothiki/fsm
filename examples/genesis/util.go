package main

import (
	"errors"
	"fmt"

	"github.com/looplab/fsm"
)

const Maxtries = 5

type eventfunc func(string, ...interface{}) error

//eventreentrant retries and event based on numretries
func eventreentrant(event eventfunc, name string, numretries int, args ...interface{}) error {
	var err error
	for i := 0; i < numretries; i++ {
		err = event(name, true, args)
		fmt.Println(err)
	}
	err = event(name, false, args)
	fmt.Println(err)
	return err
}

// eventexecutor takes fsm object and event name to execute and expecterror bool and number of retries
// in case of error expected then goes to error event and executes error event in eventreentrant function
func eventexecutor(fsm *fsm.FSM, name string, expecterror bool, numretries int, args ...interface{}) {
	var err error
	if fsm.Can(name) {
		err = fsm.Event(name, expecterror)
		if err != nil {
			fmt.Println(err)
			fmt.Println(eventreentrant(fsm.Event, "error", numretries, args))
		}
	}
	fmt.Println(fsm.Current())
}

//Newsfsm return a simple fsm object which has state transitions as explained in readme

func Newsfsm(state string) *fsm.FSM {
	return fsm.NewFSM(
		state,
		fsm.Events{
			{Name: "init", Src: []string{"idle"}, Dst: "state1"},
			{Name: "state1_2", Src: []string{"state1"}, Dst: "state2"},
			{Name: "state2_3", Src: []string{"state2"}, Dst: "state3"},
			{Name: "state3_4", Src: []string{"state3"}, Dst: "state4"},
			{Name: "close", Src: []string{"state4"}, Dst: "closed"},
			{Name: "error", Src: []string{"state1"}, Dst: "state1"},
			{Name: "error", Src: []string{"state2"}, Dst: "state2"},
			{Name: "error", Src: []string{"state3"}, Dst: "state3"},
			{Name: "error", Src: []string{"state4"}, Dst: "state4"},
		},
		fsm.Callbacks{
			"before_event": func(e *fsm.Event) {
				//fmt.Println("before_event")
			},
			"after_event": func(e *fsm.Event) {
				//fmt.Println("after_event")
			},
			"init": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				if e.Args[0].(bool) {
					e.Err = errors.New("error init")
				}
			},
			"state1_2": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error makePlan")
				}
			},
			"state2_3": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error planexecute")
				}
			},
			"state3_4": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error executeJobs")
				}
			},
			"close": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
			},
			"error": func(e *fsm.Event) {
				fmt.Println(e.FSM.Current())
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error reentered")
				}
			},
		},
	)
}

//Newsfsm return a simple fsm object which has state transitions as explained below
/*    transiiton event |--->---|                    |--->---|                       |--->---|                           |--->---|
idle---->(init)---->initcomplete---->(makePlan)---->planinit---->(planexecute)---->executeplan---->(executeJobs)----->completeexecution
										   |---<---|                    |---<---|                       |---<---|                            |---<---|
---->(close)----->closed

func Newsfsm1(state string) *fsm.FSM {
	return fsm.NewFSM(
		state,
		fsm.Events{
			{Name: "init", Src: []string{"idle"}, Dst: "initcomplete"},
			{Name: "makePlan", Src: []string{"initcomplete"}, Dst: "planinit"},
			{Name: "planexecute", Src: []string{"planinit"}, Dst: "executeplan"},
			{Name: "executeJobs", Src: []string{"executeplan"}, Dst: "completeexecution"},
			{Name: "close", Src: []string{"completeexecution"}, Dst: "closed"},
			{Name: "error", Src: []string{"initcomplete"}, Dst: "initcomplete"},
			{Name: "error", Src: []string{"planinit"}, Dst: "planinit"},
			{Name: "error", Src: []string{"executeplan"}, Dst: "executeplan"},
			{Name: "error", Src: []string{"completeexecution"}, Dst: "completeexecution"},
		},
		fsm.Callbacks{
			"before_event": func(e *fsm.Event) {
				//fmt.Println("before_event")
			},
			"after_event": func(e *fsm.Event) {
				//fmt.Println("after_event")
			},
			"init": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				if e.Args[0].(bool) {
					e.Err = errors.New("error init")
				}
			},
			"makePlan": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error makePlan")
				}
			},
			"planexecute": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error planexecute")
				}
			},
			"executeJobs": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error executeJobs")
				}
			},
			"close": func(e *fsm.Event) {
				fmt.Println(e.Event)
				fmt.Println(e.Src)
			},
			"error": func(e *fsm.Event) {
				fmt.Println(e.FSM.Current())
				// _, ok := e.Args[0].(bool)
				if e.Args[0].(bool) {
					e.Err = errors.New("error reentered")
				}
			},
		},
	)
}
*/
