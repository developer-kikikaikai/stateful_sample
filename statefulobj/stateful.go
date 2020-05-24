package statefulobj

import (
	"errors"
	"fmt"

	"github.com/bykof/stateful"
)

type SampleMachine struct {
	state stateful.State
}

type SampleArgument struct {
	history []string
}

func (sa *SampleArgument) GetHistory() string {
	ret := ""
	for _, val := range sa.history {
		ret += " " + val
	}
	return ret
}
func (sa *SampleArgument) SetHistory(now string) {
	sa.history = append(sa.history, now)
}

var (
	BEGIN = stateful.DefaultState("BEGIN")
	A     = stateful.DefaultState("A")
	B     = stateful.DefaultState("B")
)

func NewMachine() SampleMachine {
	return SampleMachine{state: BEGIN}
}

func (sm *SampleMachine) State() stateful.State {
	fmt.Printf("Get State:%s\n", sm.state.GetID())
	return sm.state
}

func (sm *SampleMachine) SetState(state stateful.State) error {
	fmt.Printf("Get State:%s\n", sm.state.GetID())
	fmt.Printf("Change from %s to %s\n", sm.state.GetID(), state.GetID())
	sm.state = state
	fmt.Printf("Current State:%s\n", sm.state.GetID())
	return nil
}

func (sm SampleMachine) ToB(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	//transition argument is in Run parameter
	param, ok := transitionArguments.(*SampleArgument)
	if !ok {
		return nil, errors.New("Error")
	}
	fmt.Printf("State change, FromAtoB, GetID=%s\n", B.GetID())
	param.SetHistory(sm.State().GetID() + "->" + B.GetID())
	return B, nil
}

func (sm SampleMachine) FromBtoA(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	//transition argument is in Run parameter
	param, ok := transitionArguments.(*SampleArgument)
	if !ok {
		return nil, errors.New("Error")
	}
	fmt.Printf("State change, FromBtoA\n")
	param.SetHistory(sm.State().GetID() + "->" + A.GetID())
	return A, nil
}

func (sm SampleMachine) FromBtoA2(transitionArguments stateful.TransitionArguments) (stateful.State, error) {
	//transition argument is in Run parameter
	param, ok := transitionArguments.(*SampleArgument)
	if !ok {
		return nil, errors.New("Error")
	}
	fmt.Printf("State change, FromBtoA2\n")
	param.SetHistory(sm.State().GetID() + "->" + A.GetID())
	return A, nil
}
