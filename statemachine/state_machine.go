package statemachine

import (
	"github.com/bykof/stateful"
)

type OrderStateMachine interface {
	Run()
}

func NewStateMachine(object *OrderState) stateful.StateMachine {
	stateMachine := stateful.StateMachine{
		StatefulObject: object,
	}
	//Add transition
	stateMachine.AddTransition(
		object.SelectProduct,
		stateful.States{BEGIN},
		stateful.States{InCart},
	)

	stateMachine.AddTransition(
		object.Order,
		stateful.States{BEGIN, InCart},
		stateful.States{Ordered},
	)

	stateMachine.AddTransition(
		object.Cancel,
		stateful.States{Ordered},
		stateful.States{Canceled},
	)

	stateMachine.AddTransition(
		object.ShopProblem,
		stateful.States{Ordered},
		stateful.States{Canceled},
	)

	stateMachine.AddTransition(
		object.Ship,
		stateful.States{Ordered},
		stateful.States{Shipped},
	)

	return stateMachine
}
