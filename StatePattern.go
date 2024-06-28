package main

import "fmt"

// state
type tvState interface {
	state()
}

// concrete implementation
type on struct{}

// implementing behaviour of ON state
func (o *on) state() {
	fmt.Println("TV is ON!")
}

type off struct{}

// implementing behaviour of OFF state
func (o *off) state() {
	fmt.Println("TV is OFF!")
}

type stateContext struct {
	currentTvState tvState
}

func getContext() *stateContext {
	return &stateContext{
		currentTvState: &off{},
	}
}

func (sc *stateContext) setState(state tvState) {
	sc.currentTvState = state

}

func (sc *stateContext) getState() {
	sc.currentTvState.state()
}

func main() {
	tvContext := getContext() // Default state is OFF
	tvContext.getState()      // Get the state as OFF
	tvContext.setState(&on{}) // Change the current state to ON
	tvContext.getState()      // Get the state as ON

}

// Comes under behaviour design pattern, when change the behaviour of the object based on specific state
// State Interface: Defines the common methods shared by all states.
// Concrete State Objects: Implement the State Interface and encapsulate logic for their specific state.
// Context Object: Holds a reference to the current state object and handles state transitions.
