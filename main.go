package main

import (
	"fmt"

	"stateful_sample/statemachine"

	"github.com/awalterschulze/gographviz"
	"github.com/bykof/stateful/statefulGraph"
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
	order := statemachine.NewOrderState(10000)
	machine := statemachine.NewStateMachine(order)

	//check graph
	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: machine}
	//_ = stateMachineGraph.DrawGraph()
	//nameが無いので追加。PRを送っています。
	_ = Draw("order", stateMachineGraph)

	//注文してみる
	var err error
	product1 := statemachine.Product{1000}
	err = machine.Run(order.SelectProduct, &product1)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	//2重の注文はできない
	err = machine.Run(order.SelectProduct, &product1)
	if err == nil {
		fmt.Printf("Error")
		return
	} else {
		fmt.Printf("You can't send multiple requests: %s\n", err.Error())
	}
	//注文
	err = machine.Run(order.Order, &product1)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())
	//Cancelする
	err = machine.Run(order.Cancel, &product1)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())
	//もう一度注文
	err = machine.Run(order.Order, &product1)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())
	//発送する
	err = machine.Run(order.Ship, &product1)
	if err != nil {
		fmt.Printf("Error")
		return
	}
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())
	//高いものは注文できない
	product2 := statemachine.Product{9500}
	err = machine.Run(order.Order, &product2)
	if err == nil {
		fmt.Printf("Error")
		return
	}
	fmt.Printf("%s\n", err.Error())
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())

	_ = machine.Run(order.Order, &product1)
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())

	_ = machine.Run(order.ShopProblem, &product1)
	fmt.Printf("  ##Deposit: %d\n", order.Deposit())
}
