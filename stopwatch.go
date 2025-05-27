// Package stopwatch provides a simple
// timer implementation in Go for measuring
// time in ticks
package stopwatch

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Stopwatch is a representation of the timer for Ebitengine
// it works by counting ticks
type Stopwatch struct {
	currentTicks int
	maxTicks     int
	isActive     bool
}

// NewStopwatch is a factory method which creates and returns
// a new Stopwatch instance.
func NewStopwatch(d *time.Duration) *Stopwatch {
	return &Stopwatch{
		currentTicks: 0,
		maxTicks:     int(d.Milliseconds()) * ebiten.TPS() / 1000,
		isActive:     false,
	}
}

// Begins a new stopwatch, or restarts a paused instance.
func (s *Stopwatch) Start() {
	s.isActive = true
}

// Stops or pauses a timer instance
func (s *Stopwatch) Stop() {
	s.isActive = false
}

// Update increments the timer instance.
func (s *Stopwatch) Update() {
	if s.isActive && s.currentTicks < s.maxTicks {
		s.currentTicks++
	}
}

// Resets the timer back to zero.
func (s *Stopwatch) Reset() {
	s.currentTicks = 0
}

// Checks to see if the stopwatch is finished.
// Future changes
func (s *Stopwatch) IsDone() bool {
	return s.maxTicks <= s.currentTicks
}
