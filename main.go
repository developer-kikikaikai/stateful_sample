package main

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
	"github.com/bykof/stateful"
	"github.com/bykof/stateful/statefulGraph"
	. "stateful_sample/statefulobj"
)

func Draw(name string, smg statefulGraph.StateMachineGraph) error {
	var err error
	graph := gographviz.NewGraph()

	err = graph.SetDir(true)
	if err != nil {
		return err
	}

	err = graph.SetName(name)
	if err != nil {
		return err
	}

	err = smg.DrawStates(graph)
	if err != nil {
		return err
	}

	err = smg.DrawEdges(graph)
	if err != nil {
		return err
	}

	fmt.Println(graph.String())

	return nil
}

func main() {
	machine := NewMachine()
	stateMachine := stateful.StateMachine{
		StatefulObject: &machine,
	}
	//Add transition
	stateMachine.AddTransition(
		machine.ToB,
		stateful.States{BEGIN, A}, //src
		stateful.States{B},        //dist
	)

	stateMachine.AddTransition(
		machine.FromBtoA,
		stateful.States{B}, //src
		stateful.States{A}, //dist
	)
	stateMachine.AddTransition(
		machine.FromBtoA2,
		stateful.States{B}, //src
		stateful.States{A}, //dist
	)

	//check automaton
	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: stateMachine}
	_ = stateMachineGraph.DrawGraph()
	_ = Draw("sample", stateMachineGraph)

	arg := SampleArgument{}
	fmt.Printf("[main] expect: failed!!!\n")
	err := stateMachine.Run(
		machine.FromBtoA,
		&arg,
	)
	if err != nil {
		fmt.Printf("###Error!!%s\n", err.Error())
	}
	fmt.Printf("[main] Run with ToB!!!\n")
	err = stateMachine.Run(
		machine.ToB,
		&arg,
	)
	if err != nil {
		fmt.Printf("###Error!!%s\n", err.Error())
	}

	fmt.Printf("[main] Run with FromBtoA!!!\n")
	err = stateMachine.Run(
		machine.FromBtoA,
		&arg,
	)
	if err != nil {
		fmt.Printf("###Error!!%s\n", err.Error())
	}
	fmt.Printf("[main] Run with FromAtoB!!!\n")
	err = stateMachine.Run(
		machine.ToB,
		&arg,
	)
	if err != nil {
		fmt.Printf("###Error!!%s\n", err.Error())
	}
	fmt.Printf("[main] Run!!!\n")
	err = stateMachine.Run(
		machine.FromBtoA2,
		&arg,
	)
	if err != nil {
		fmt.Printf("###Error!!%s\n", err.Error())
	}

	fmt.Printf("[main] History:%s\n", arg.GetHistory())
}
