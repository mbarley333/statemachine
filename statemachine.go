package statemachine

import "errors"

var ErrEventRejected = errors.New("event rejected")

// event is the action the could force a change
type EventType int

// result of event
type StateType int

const (
	StateTypeDefault StateType = 0

	EventTypeNoOp = 0
)

type Events map[EventType]StateType

// Events that are ok for a given state
type State struct {
	Events Events
}

type Machine struct {
	CurrentState        StateType
	StateTypeToStateMap map[StateType]State
}

func (m *Machine) GetNextState(event EventType) (StateType, error) {

	state, okState := m.StateTypeToStateMap[m.CurrentState]
	if okState {
		next, okEvent := state.Events[event]
		if okEvent {
			return next, nil
		}
	}
	return StateTypeDefault, ErrEventRejected
}
