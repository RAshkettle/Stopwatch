package stopwatch

import (
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// TestNewStopwatch tests the factory method for creating a new Stopwatch
func TestNewStopwatch(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected struct {
			currentTicks int
			maxTicks     int
			isActive     bool
		}
	}{
		{
			name:     "1 second duration",
			duration: time.Second,
			expected: struct {
				currentTicks int
				maxTicks     int
				isActive     bool
			}{
				currentTicks: 0,
				maxTicks:     ebiten.TPS(),
				isActive:     false,
			},
		},
		{
			name:     "500 milliseconds duration",
			duration: 500 * time.Millisecond,
			expected: struct {
				currentTicks int
				maxTicks     int
				isActive     bool
			}{
				currentTicks: 0,
				maxTicks:     ebiten.TPS() / 2,
				isActive:     false,
			},
		},
		{
			name:     "2 seconds duration",
			duration: 2 * time.Second,
			expected: struct {
				currentTicks int
				maxTicks     int
				isActive     bool
			}{
				currentTicks: 0,
				maxTicks:     2 * ebiten.TPS(),
				isActive:     false,
			},
		},
		{
			name:     "zero duration",
			duration: 0,
			expected: struct {
				currentTicks int
				maxTicks     int
				isActive     bool
			}{
				currentTicks: 0,
				maxTicks:     0,
				isActive:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := NewStopwatch(tt.duration)

			if sw == nil {
				t.Fatal("NewStopwatch returned nil")
			}

			if sw.currentTicks != tt.expected.currentTicks {
				t.Errorf("currentTicks = %d, want %d", sw.currentTicks, tt.expected.currentTicks)
			}

			if sw.maxTicks != tt.expected.maxTicks {
				t.Errorf("maxTicks = %d, want %d", sw.maxTicks, tt.expected.maxTicks)
			}

			if sw.isActive != tt.expected.isActive {
				t.Errorf("isActive = %t, want %t", sw.isActive, tt.expected.isActive)
			}
		})
	}
}

// TestStopwatchStart tests the Start method
func TestStopwatchStart(t *testing.T) {
	duration := time.Second
	sw := NewStopwatch(duration)

	// Initially should be inactive
	if sw.isActive {
		t.Error("Stopwatch should be inactive initially")
	}

	// Start the stopwatch
	sw.Start()

	if !sw.isActive {
		t.Error("Stopwatch should be active after calling Start()")
	}

	// Starting an already active stopwatch should keep it active
	sw.Start()
	if !sw.isActive {
		t.Error("Stopwatch should remain active after calling Start() again")
	}
}

// TestStopwatchStop tests the Stop method
func TestStopwatchStop(t *testing.T) {
	duration := time.Second
	sw := NewStopwatch(duration)

	// Start the stopwatch first
	sw.Start()
	if !sw.isActive {
		t.Error("Stopwatch should be active after Start()")
	}

	// Stop the stopwatch
	sw.Stop()
	if sw.isActive {
		t.Error("Stopwatch should be inactive after calling Stop()")
	}

	// Stopping an already inactive stopwatch should keep it inactive
	sw.Stop()
	if sw.isActive {
		t.Error("Stopwatch should remain inactive after calling Stop() again")
	}
}

// TestStopwatchUpdate tests the Update method
func TestStopwatchUpdate(t *testing.T) {
	duration := 100 * time.Millisecond
	sw := NewStopwatch(duration)

	// Test update when stopwatch is inactive
	initialTicks := sw.currentTicks
	sw.Update()
	if sw.currentTicks != initialTicks {
		t.Error("Update() should not increment ticks when stopwatch is inactive")
	}

	// Test update when stopwatch is active
	sw.Start()
	for i := 0; i < 5; i++ {
		sw.Update()
	}
	if sw.currentTicks != 5 {
		t.Errorf("Expected currentTicks to be 5, got %d", sw.currentTicks)
	}

	// Test that update stops incrementing after reaching maxTicks
	sw.currentTicks = sw.maxTicks
	beforeUpdate := sw.currentTicks
	sw.Update()
	if sw.currentTicks != beforeUpdate {
		t.Error("Update() should not increment beyond maxTicks")
	}
}

// TestStopwatchReset tests the Reset method
func TestStopwatchReset(t *testing.T) {
	duration := time.Second
	sw := NewStopwatch(duration)

	// Start and update the stopwatch
	sw.Start()
	for i := 0; i < 10; i++ {
		sw.Update()
	}

	if sw.currentTicks == 0 {
		t.Error("currentTicks should be greater than 0 after updates")
	}

	// Reset the stopwatch
	sw.Reset()

	if sw.currentTicks != 0 {
		t.Errorf("currentTicks should be 0 after Reset(), got %d", sw.currentTicks)
	}

	// Verify that isActive state is preserved after reset
	if !sw.isActive {
		t.Error("isActive state should be preserved after Reset()")
	}
}

// TestStopwatchIsDone tests the IsDone method
func TestStopwatchIsDone(t *testing.T) {
	tests := []struct {
		name         string
		currentTicks int
		maxTicks     int
		expected     bool
	}{
		{
			name:         "not done - currentTicks less than maxTicks",
			currentTicks: 5,
			maxTicks:     10,
			expected:     false,
		},
		{
			name:         "done - currentTicks equals maxTicks",
			currentTicks: 10,
			maxTicks:     10,
			expected:     true,
		},
		{
			name:         "done - currentTicks greater than maxTicks",
			currentTicks: 15,
			maxTicks:     10,
			expected:     true,
		},
		{
			name:         "edge case - both zero",
			currentTicks: 0,
			maxTicks:     0,
			expected:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration := time.Second
			sw := NewStopwatch(duration)
			sw.currentTicks = tt.currentTicks
			sw.maxTicks = tt.maxTicks

			result := sw.IsDone()
			if result != tt.expected {
				t.Errorf("IsDone() = %t, want %t", result, tt.expected)
			}
		})
	}
}

// TestStopwatchLifecycle tests the complete lifecycle of a stopwatch
func TestStopwatchLifecycle(t *testing.T) {
	duration := 50 * time.Millisecond // Short duration for quick test
	sw := NewStopwatch(duration)

	// Initial state
	if sw.IsDone() {
		t.Error("Stopwatch should not be done initially")
	}
	if sw.isActive {
		t.Error("Stopwatch should not be active initially")
	}

	// Start and run until completion
	sw.Start()
	if !sw.isActive {
		t.Error("Stopwatch should be active after Start()")
	}

	// Update until done
	for !sw.IsDone() {
		sw.Update()
	}

	if !sw.IsDone() {
		t.Error("Stopwatch should be done after reaching maxTicks")
	}

	// Reset and verify state
	sw.Reset()
	if sw.IsDone() {
		t.Error("Stopwatch should not be done after Reset()")
	}
	if sw.currentTicks != 0 {
		t.Error("currentTicks should be 0 after Reset()")
	}
	if !sw.isActive {
		t.Error("Stopwatch should remain active after Reset()")
	}
}

// TestStopwatchStartStopCycle tests starting and stopping multiple times
func TestStopwatchStartStopCycle(t *testing.T) {
	duration := time.Second
	sw := NewStopwatch(duration)

	// Test multiple start/stop cycles
	for i := 0; i < 3; i++ {
		sw.Start()
		if !sw.isActive {
			t.Errorf("Stopwatch should be active after Start() cycle %d", i+1)
		}

		sw.Stop()
		if sw.isActive {
			t.Errorf("Stopwatch should be inactive after Stop() cycle %d", i+1)
		}
	}
}

// TestStopwatchEdgeCases tests edge cases and boundary conditions
func TestStopwatchEdgeCases(t *testing.T) {
	t.Run("very small duration", func(t *testing.T) {
		duration := time.Nanosecond
		sw := NewStopwatch(duration)
		if sw.maxTicks < 0 {
			t.Error("maxTicks should not be negative for very small durations")
		}
	})
	
	t.Run("very large duration", func(t *testing.T) {
		duration := 24 * time.Hour
		sw := NewStopwatch(duration)
		expectedMaxTicks := int(duration.Milliseconds()) * ebiten.TPS() / 1000
		if sw.maxTicks != expectedMaxTicks {
			t.Errorf("maxTicks = %d, want %d for large duration", sw.maxTicks, expectedMaxTicks)
		}
	})

	t.Run("update after done", func(t *testing.T) {
		duration := time.Millisecond
		sw := NewStopwatch(duration)
		sw.Start()

		// Run until done
		for !sw.IsDone() {
			sw.Update()
		}

		ticksWhenDone := sw.currentTicks

		// Update again - should not increment
		sw.Update()
		if sw.currentTicks != ticksWhenDone {
			t.Error("Update() should not increment ticks after stopwatch is done")
		}
	})
}

// TestStopwatchIsRunning tests the IsRunning method
func TestStopwatchIsRunning(t *testing.T) {
	tests := []struct {
		name         string
		setupFunc    func(*Stopwatch)
		expected     bool
		description  string
	}{
		{
			name: "initially not running",
			setupFunc: func(sw *Stopwatch) {
				// Do nothing - test initial state
			},
			expected:    false,
			description: "new stopwatch should not be running initially",
		},
		{
			name: "running after start",
			setupFunc: func(sw *Stopwatch) {
				sw.Start()
			},
			expected:    true,
			description: "stopwatch should be running after Start()",
		},
		{
			name: "not running after stop",
			setupFunc: func(sw *Stopwatch) {
				sw.Start()
				sw.Stop()
			},
			expected:    false,
			description: "stopwatch should not be running after Stop()",
		},
		{
			name: "not running when done but active",
			setupFunc: func(sw *Stopwatch) {
				sw.Start()
				// Set to completed state
				sw.currentTicks = sw.maxTicks
			},
			expected:    false,
			description: "stopwatch should not be running when done, even if active",
		},
		{
			name: "not running when done and stopped",
			setupFunc: func(sw *Stopwatch) {
				sw.Start()
				// Set to completed state
				sw.currentTicks = sw.maxTicks
				sw.Stop()
			},
			expected:    false,
			description: "stopwatch should not be running when done and stopped",
		},
		{
			name: "running after reset",
			setupFunc: func(sw *Stopwatch) {
				sw.Start()
				for i := 0; i < 5; i++ {
					sw.Update()
				}
				sw.Reset()
			},
			expected:    true,
			description: "stopwatch should be running after reset if it was active",
		},
		{
			name: "not running with zero duration",
			setupFunc: func(sw *Stopwatch) {
				sw.maxTicks = 0 // Zero duration
				sw.Start()
			},
			expected:    false,
			description: "stopwatch with zero duration should not be running even when started",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration := time.Second
			sw := NewStopwatch(duration)

			// Apply the setup function
			tt.setupFunc(sw)

			result := sw.IsRunning()
			if result != tt.expected {
				t.Errorf("IsRunning() = %t, want %t - %s", result, tt.expected, tt.description)
			}
		})
	}
}

// TestStopwatchIsRunningStateTransitions tests IsRunning during various state transitions
func TestStopwatchIsRunningStateTransitions(t *testing.T) {
	duration := 100 * time.Millisecond
	sw := NewStopwatch(duration)

	// Initial state - should not be running
	if sw.IsRunning() {
		t.Error("Stopwatch should not be running initially")
	}

	// Start - should be running
	sw.Start()
	if !sw.IsRunning() {
		t.Error("Stopwatch should be running after Start()")
	}

	// Update a few times - should still be running
	for i := 0; i < 3; i++ {
		sw.Update()
		if !sw.IsRunning() {
			t.Errorf("Stopwatch should still be running after %d updates", i+1)
		}
	}

	// Stop - should not be running
	sw.Stop()
	if sw.IsRunning() {
		t.Error("Stopwatch should not be running after Stop()")
	}

	// Start again - should be running
	sw.Start()
	if !sw.IsRunning() {
		t.Error("Stopwatch should be running after restarting")
	}

	// Update until done - should not be running when done
	for !sw.IsDone() {
		sw.Update()
	}
	if sw.IsRunning() {
		t.Error("Stopwatch should not be running when done, even if active")
	}

	// Reset - should be running again (since it was active)
	sw.Reset()
	if !sw.IsRunning() {
		t.Error("Stopwatch should be running after Reset() when it was previously active")
	}
}

// TestStopwatchIsRunningEdgeCases tests edge cases for IsRunning
func TestStopwatchIsRunningEdgeCases(t *testing.T) {
	t.Run("zero duration stopwatch", func(t *testing.T) {
		duration := time.Duration(0)
		sw := NewStopwatch(duration)

		// Should not be running initially
		if sw.IsRunning() {
			t.Error("Zero duration stopwatch should not be running initially")
		}

		// Should not be running even after start (since it's immediately done)
		sw.Start()
		if sw.IsRunning() {
			t.Error("Zero duration stopwatch should not be running even after Start()")
		}
	})

	t.Run("very small duration", func(t *testing.T) {
		duration := time.Nanosecond
		sw := NewStopwatch(duration)

		sw.Start()
		// Should be running (since maxTicks would be 0 but IsDone() returns maxTicks <= currentTicks)
		if sw.maxTicks == 0 && sw.IsRunning() {
			t.Error("Very small duration stopwatch should not be running if maxTicks is 0")
		}
	})

	t.Run("manually manipulated state", func(t *testing.T) {
		duration := time.Second
		sw := NewStopwatch(duration)

		// Manually set to active but done
		sw.isActive = true
		sw.currentTicks = sw.maxTicks + 1

		if sw.IsRunning() {
			t.Error("Stopwatch should not be running when manually set to done state")
		}

		// Reset current ticks but keep active
		sw.currentTicks = 0

		if !sw.IsRunning() {
			t.Error("Stopwatch should be running when active and not done")
		}
	})
}

// TestStopwatchIsRunningConsistency tests consistency between IsRunning and other methods
func TestStopwatchIsRunningConsistency(t *testing.T) {
	duration := time.Second
	sw := NewStopwatch(duration)

	// Test consistency throughout lifecycle
	states := []struct {
		action      func()
		description string
	}{
		{func() {}, "initial state"},
		{func() { sw.Start() }, "after start"},
		{func() { sw.Update() }, "after first update"},
		{func() { 
			for i := 0; i < 10; i++ {
				sw.Update()
			}
		}, "after multiple updates"},
		{func() { sw.Stop() }, "after stop"},
		{func() { sw.Start() }, "after restart"},
		{func() { sw.Reset() }, "after reset"},
		{func() {
			for !sw.IsDone() {
				sw.Update()
			}
		}, "after completion"},
	}

	for _, state := range states {
		state.action()

		isRunning := sw.IsRunning()
		isActive := sw.isActive
		isDone := sw.IsDone()

		// IsRunning should be true only when active AND not done
		expectedRunning := isActive && !isDone

		if isRunning != expectedRunning {
			t.Errorf("%s: IsRunning() = %t, but expected %t (active=%t, done=%t)", 
				state.description, isRunning, expectedRunning, isActive, isDone)
		}
	}
}

// BenchmarkStopwatchUpdate benchmarks the Update method
func BenchmarkStopwatchUpdate(b *testing.B) {
	duration := time.Second
	sw := NewStopwatch(duration)
	sw.Start()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sw.Update()
		if sw.IsDone() {
			sw.Reset()
		}
	}
}

// BenchmarkStopwatchIsDone benchmarks the IsDone method
func BenchmarkStopwatchIsDone(b *testing.B) {
	duration := time.Second
	sw := NewStopwatch(duration)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sw.IsDone()
	}
}

// BenchmarkStopwatchIsRunning benchmarks the IsRunning method
func BenchmarkStopwatchIsRunning(b *testing.B) {
	duration := time.Second
	sw := NewStopwatch(duration)
	sw.Start()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sw.IsRunning()
	}
}
