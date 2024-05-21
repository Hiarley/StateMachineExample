package main

import "StateMacheExample/statemachine"

func main() {
	received := statemachine.Event{Name: "received", Code: "RECEIVED"}
	matched := statemachine.Event{Name: "matched", Code: "MATCHED"}
	discarded := statemachine.Event{Name: "discarded", Code: "DISCARDED"}
	consolidated := statemachine.Event{Name: "consolidated", Code: "CONSOLIDATED"}
	unconsolidated := statemachine.Event{Name: "unconsolidated", Code: "UNCONSOLIDATED"}

	logReceivedCmd := statemachine.Command{Name: "State: Received", Code: "LOG_RECEIVED"}
	logProcessMatchedCmd := statemachine.Command{Name: "Process Match", Code: "PROCESS_MATCHED"}
	logMatchedCmd := statemachine.Command{Name: "State: Matched", Code: "LOG_MATCHED"}
	logProcessConsolidatedCmd := statemachine.Command{Name: "Process Consolidated", Code: "PROCESS_CONSOLIDATED"}

	logProcessDiscardedCmd := statemachine.Command{Name: "Process Discarded", Code: "PROCESS_DISCARDED"}
	logDiscardedCmd := statemachine.Command{Name: "State: Discarded", Code: "LOG_DISCARDED"}
	logConsolidatedCmd := statemachine.Command{Name: "State: Consolidated", Code: "LOG_CONSOLIDATED"}
	logUnconsolidatedCmd := statemachine.Command{Name: "State: Unconsolidated", Code: "LOG_UNCONSOLIDATED"}

	receivedState := statemachine.NewState("received")
	matchedState := statemachine.NewState("matched")
	discardedState := statemachine.NewState("discarded")
	consolidatedState := statemachine.NewState("consolidated")
	unconsolidatedState := statemachine.NewState("unconsolidated")

	receivedState.AddTransition(matched, matchedState)
	receivedState.AddTransition(discarded, discardedState)

	receivedState.AddAction(logReceivedCmd)

	matchedState.AddTransition(consolidated, consolidatedState)
	matchedState.AddTransition(unconsolidated, unconsolidatedState)

	matchedState.AddAction(logProcessMatchedCmd)
	matchedState.AddAction(logMatchedCmd)

	discardedState.AddAction(logProcessDiscardedCmd)
	discardedState.AddAction(logDiscardedCmd)

	consolidatedState.AddAction(logProcessConsolidatedCmd)
	consolidatedState.AddAction(logConsolidatedCmd)

	unconsolidatedState.AddAction(logUnconsolidatedCmd)

	machine := statemachine.NewStateMachine(receivedState)
	machine.AddResetEvent(received)

	// Test the state machine
	machine.HandleEvent(received)
	//machine.HandleEvent(matched)
	machine.HandleEvent(discarded)
	//machine.HandleEvent(consolidated)
	//machine.HandleEvent(unconsolidated)
	//machine.HandleEvent(received)
}
