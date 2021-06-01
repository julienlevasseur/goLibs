package stopwatch_test

import (
	"testing"

	"github.com/julienlevasseur/goLibs/stopwatch"
	"github.com/stretchr/testify/require"
)

func TestStopwatch(t *testing.T) {
	sw := stopwatch.New(0, false)
	sw.Start()
	require.Equal(t, true, sw.Active())
	sw.Stop()
	require.Equal(t, false, sw.Active())
	sw.Reset(0, true)
	require.Equal(t, true, sw.Active())
	sw.Reset(0, false)
	require.Equal(t, false, sw.Active())
	require.NotEqual(t, 0, sw.ElapsedTime())
}
