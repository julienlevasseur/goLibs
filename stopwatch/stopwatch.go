package stopwatch

import (
	"fmt"
	"time"
)

// Stopwatch implements the stopwatch functionality
type Stopwatch struct {
	start, stop time.Time
}

// SW is the global Stopwatch instance
var SW *Stopwatch

// New creates a new stopwatch with starting time offset by
// a user defined value. Negative offsets result in a countdown
// prior to the start of the stopwatch.
func New(offset time.Duration, active bool) *Stopwatch {
	var sw Stopwatch
	sw.Reset(offset, active)
	return &sw
}

// Reset allows the re-use of a Stopwatch instead of creating
// a new one.
func (s *Stopwatch) Reset(offset time.Duration, active bool) {
	now := time.Now()
	s.start = now.Add(-offset)
	if active {
		s.stop = time.Time{}
	} else {
		s.stop = now
	}
}

// Active returns true if the stopwatch is active (counting up)
func (s *Stopwatch) Active() bool {
	return s.stop.IsZero()
}

// Stop makes the stopwatch stop counting up
func (s *Stopwatch) Stop() {
	if s.Active() {
		s.stop = time.Now()
	}
}

// Run initiates, or resumes the counting up process
func (s *Stopwatch) Run() {
	if !s.Active() {
		diff := time.Now().Sub(s.stop)
		s.start = s.start.Add(diff)
		s.stop = time.Time{}
	}
}

// ElapsedTime is the time the stopwatch has been active
func (s *Stopwatch) ElapsedTime() time.Duration {
	if s.Active() {
		return time.Since(s.start)
	}
	return s.stop.Sub(s.start)
}

// String representation of a single Stopwatch instance.
func (s *Stopwatch) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]",
		s.start.Format(time.Stamp), time.Now().Format(time.Stamp), s.ElapsedTime())
}
