package replog

import (
	"fmt"

	"github.com/rotisserie/eris"
	"github.com/sashajdn/replog/pkg/log"
)

var ErrInvalidStateTransition = eris.New("invalid state transition")
var ErrTransitionToSameState = eris.New("transition to same state")

var validTransitions = map[state]map[state]struct{}{
	ready:      {collecting: struct{}{}},
	collecting: {cancelled: struct{}{}, finished: struct{}{}, collecting: struct{}{}},
	finished:   {ready: struct{}{}},
	cancelled:  {ready: struct{}{}},
}

type state uint16

const (
	ready state = iota
	collecting
	finished
	cancelled
)

func (s state) String() string {
	switch s {
	case cancelled:
		return "cancelled"
	case collecting:
		return "collecting"
	case finished:
		return "finished"
	case ready:
		return "ready"
	default:
		// We should never hit this case.
		panic(fmt.Sprintf("unexpected replog.state: %#v", s))
	}
}

func (s *state) transition(to state) error {
	validTransitions, ok := validTransitions[*s]
	if !ok {
		panic("we are currently in an invalid state")
	}

	if _, ok := validTransitions[to]; !ok {
		return eris.Wrap(ErrInvalidStateTransition, fmt.Sprintf("invalid state transition from %v to %v", s, to))
	}

	*s = to
	return nil
}

func NewStateMachine(logger *log.SugaredLogger) *stateMachine {
	return &stateMachine{
		currentState: ready,
		logger:       logger,
	}
}

type stateMachine struct {
	currentState state
	logger       *log.SugaredLogger
}

func (s *stateMachine) CurrentState() state {
	return s.currentState
}

func (s *stateMachine) SetReady() error {
	return s.transition(ready)
}

func (s *stateMachine) SetCancelled() error {
	return s.transition(cancelled)
}

func (s *stateMachine) SetFinished() error {
	return s.transition(finished)
}

func (s *stateMachine) SetCollecting() error {
	return s.transition(collecting)
}

func (s *stateMachine) transition(to state) error {
	if s.currentState == to {
		return ErrTransitionToSameState
	}

	s.logger.Infof("transitioning state machine from %s to %s", s.currentState, to)
	return s.currentState.transition(to)
}
