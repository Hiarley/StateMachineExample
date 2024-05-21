package statemachine

import "fmt"

type StateMachine struct {
	InitialState *State
	CurrentState *State
	ResetEvents  map[string]bool
}

func NewStateMachine(initialState *State) *StateMachine {
	return &StateMachine{
		InitialState: initialState,
		CurrentState: initialState,
		ResetEvents:  make(map[string]bool),
	}
}

func (sm *StateMachine) AddResetEvent(event Event) {
	sm.ResetEvents[event.Code] = true
}

func (sm *StateMachine) HandleEvent(event Event) {
	if sm.ResetEvents[event.Code] {
		fmt.Println("Initial State: ", sm.InitialState.Name)
		sm.CurrentState = sm.InitialState
		for _, action := range sm.CurrentState.Actions {
			sm.ExecuteCommand(action)
		}
	} else {
		if nextState, ok := sm.CurrentState.Transitions[event.Code]; ok {
			sm.CurrentState = nextState
			for _, action := range sm.CurrentState.Actions {
				sm.ExecuteCommand(action)
			}
		}
	}
}

func (sm *StateMachine) ExecuteCommand(command Command) {
	fmt.Println("Executing command:", command.Name)
}
