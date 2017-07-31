package main

import (
	"log"
	"fmt"
	"github.com/ryanfaerman/fsm"
)

type Stream struct {
	State fsm.State

	// our machine cache
	machine *fsm.Machine
}

func newRules () *fsm.Ruleset {
	// Establish some rules for our FSM
	rules := fsm.Ruleset{}
	rules.AddTransition(fsm.T{"created", "active"})
	rules.AddTransition(fsm.T{"active", "interrupted"})
	rules.AddTransition(fsm.T{"interrupted", "active"})
	rules.AddTransition(fsm.T{"interrupted", "finished"})

	return &rules
}

// Add methods to comply with the fsm.Stater interface
func (t *Stream) CurrentState() fsm.State { return t.State }
func (t *Stream) SetState(s fsm.State)    { t.State = s }

// A helpful function that lets us apply arbitrary rulesets to this
// instances state machine without reallocating the machine. While not
// required, it's something I like to have.
func (t *Stream) Apply(r *fsm.Ruleset) *fsm.Machine {
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}

	t.machine.Rules = r
	return t.machine
}

func main() {
	var err error

	some_stream := Stream{State: "created"} // Our subject
	fmt.Println(some_stream)

	rules := newRules()

	err = some_stream.Apply(rules).Transition("active")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status", some_stream)
}