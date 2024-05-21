package statemachine

type Event struct {
	Name string
	Code string
}

type Command struct {
	Name string
	Code string
}

type State struct {
	Name        string
	Transitions map[string]*State
	Actions     []Command
}

func NewState(name string) *State {
	return &State{
		Name:        name,
		Transitions: make(map[string]*State),
		Actions:     []Command{},
	}
}

func (s *State) AddTransition(event Event, state *State) {
	s.Transitions[event.Code] = state
}

func (s *State) AddAction(command Command) {
	s.Actions = append(s.Actions, command)
}
