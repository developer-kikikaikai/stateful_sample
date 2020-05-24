package statemachine

import (
	"errors"
	"fmt"

	"github.com/bykof/stateful"
)

type OrderState struct {
	state   stateful.State
	deposit int
}

func NewOrderState(deposit int) *OrderState {
	return &OrderState{
		state:   BEGIN,
		deposit: deposit,
	}
}

func (s *OrderState) Deposit() int {
	return s.deposit
}

func (s *OrderState) State() stateful.State {
	return s.state
}

//this will be called after calling a transition function
func (s *OrderState) SetState(state stateful.State) error {
	s.state = state
	return nil
}

/*trasition*/
func (s *OrderState) SelectProduct(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Printf("The product is in your cart.\n")
	return InCart, nil
}

func (s *OrderState) Order(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	//transition argument is in Run parameter
	p, ok := transitionArguments.(*Product)
	if !ok {
		return nil, errors.New("Invalid argument")
	}

	if s.deposit < p.Fee {
		return nil, errors.New("The deposit is not enough")
	}
	s.deposit -= p.Fee
	fmt.Printf("Ordered!!\n")
	return Ordered, nil
}

func (s *OrderState) Cancel(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Printf("User cancels order.\n")
	//transition argument is in Run parameter
	p, ok := transitionArguments.(*Product)
	if !ok {
		return nil, errors.New("Invalid argument")
	}
	s.deposit += p.Fee
	return Canceled, nil
}

func (s *OrderState) ShopProblem(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Printf("Sorry, the product shop has some problem, cancel order.\n")
	//transition argument is in Run parameter
	p, ok := transitionArguments.(*Product)
	if !ok {
		return nil, errors.New("Invalid argument")
	}
	s.deposit += p.Fee
	return Canceled, nil
}

func (s *OrderState) Ship(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	fmt.Printf("Ship the pruduct to user's address. Please wait 2-3 days.\n")
	return Shipped, nil
}
