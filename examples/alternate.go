// +build ignore

package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "stage1", Src: []string{"idle"}, Dst: "success"},
			{Name: "stage1", Src: []string{"idle"}, Dst: "failure"},
			{Name: "stage2", Src: []string{"error", "updating"}, Dst: "success"},
		},
		fsm.Callbacks{
			"stage1": func(e *fsm.Event) {
				e.Dst = "success"
				fmt.Println("after_scan: " + e.FSM.Current())

				//e.Err = errors.New("type")
			},
			"stage2": func(e *fsm.Event) {
				fmt.Println("working: " + e.FSM.Current())
			},
			"situation": func(e *fsm.Event) {
				fmt.Println("situation: " + e.FSM.Current())
			},
		},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event("stage1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + fsm.Current())

	err = fsm.Event("stage2")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("2:" + fsm.Current())
	//
	// err = fsm.Event("situation")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("3:" + fsm.Current())
	//
	// err = fsm.Event("finish")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println("4:" + fsm.Current())

}
