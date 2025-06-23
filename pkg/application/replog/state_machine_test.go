package replog

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestStateMachine(t *testing.T) {
	t.Parallel()

	logger, _ := zap.NewDevelopment()

	sm := NewStateMachine(logger.Sugar())
	require.Equal(t, ready, sm.CurrentState())

	// Start, finished, ready path.
	require.NoError(t, sm.SetCollecting())
	require.NoError(t, sm.SetFinished())
	require.NoError(t, sm.SetReady())

	// Start, cancelled, ready path.
	require.NoError(t, sm.SetCollecting())
	require.NoError(t, sm.SetCancelled())
	require.NoError(t, sm.SetReady())

	// Invalid state transitions.
	require.ErrorIs(t, sm.SetReady(), ErrTransitionToSameState)
	require.ErrorIs(t, sm.SetFinished(), ErrInvalidStateTransition)

	// Validate after invalid state transitions we are the correct point.
	require.Equal(t, ready, sm.CurrentState())

	// Backwards in time test.
	require.NoError(t, sm.SetCollecting())
	require.ErrorIs(t, sm.SetReady(), ErrInvalidStateTransition)
	require.Equal(t, collecting, sm.CurrentState())
}
