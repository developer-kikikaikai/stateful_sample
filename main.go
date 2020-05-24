package main

import (
	"fmt"

	"stateful_sample/statemachine"

	"github.com/bykof/stateful/statefulGraph"
)

func main() {
	order := statemachine.NewOrderState(10000)
	machine := statemachine.NewStateMachine(order)

	//check graph
	stateMachineGraph := statefulGraph.StateMachineGraph{StateMachine: machine}
	_ = stateMachineGraph.DrawGraphWithName("order")

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
