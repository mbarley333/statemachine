package statemachine_test

import (
	"statemachine"
	"testing"
)

const (
	StateTypeOpcode statemachine.StateType = iota + 1
	StateTypeOperand
)

const (
	EventTypeOpcode = iota + 1
	EventTypeOperand
)

func TestTransition(t *testing.T) {
	t.Parallel()

	// struct to hold state
	s := newStatemachine()

	want := StateTypeOpcode

	got, err := s.GetNextState(EventTypeOpcode)
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}

}

func newStatemachine() *statemachine.Machine {
	return &statemachine.Machine{
		StateTypeToStateMap: map[statemachine.StateType]statemachine.State{
			statemachine.StateTypeDefault: statemachine.State{
				Events: statemachine.Events{
					EventTypeOpcode: StateTypeOpcode,
				},
			},
			StateTypeOpcode: statemachine.State{

				Events: statemachine.Events{
					EventTypeOperand: StateTypeOperand,
				},
			},
		},
	}
}
