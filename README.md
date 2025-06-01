# Stopwatch

Go/Ebitengine Timer for games

Usage is pretty basic, just import into your application with
`go get github.com/RAshkettle/Stopwatch`

## Methods

- `NewStopwatch(duration)` - Create a new stopwatch with the specified duration
- `Start()` - Start or resume the stopwatch
- `Stop()` - Stop or pause the stopwatch
- `Update()` - Increment the timer (call this every frame)
- `Reset()` - Reset the timer back to zero
- `IsDone()` - Check if the stopwatch has finished counting
- `IsRunning()` - Check if the stopwatch is currently running (active and not done)

### Example:

```go
package main

import (
	"fmt"
	"time"

	stopwatch "github.com/RAshkettle/Stopwatch"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	timer *stopwatch.Stopwatch
}

func main() {
	timer := stopwatch.NewStopwatch(200 * time.Millisecond)
	g := &Game{
		timer: timer,
	}
	timer.Start()
	ebiten.RunGame(g)
}

// Ebigengine Game Interface
func (g *Game) Draw(screen *ebiten.Image) {
	g.timer.Update()

	if g.timer.IsRunning() {
		// Timer is actively counting down
		fmt.Println("Timer is running...")
	}

	if g.timer.IsDone() {
		fmt.Println("DING!")
		g.timer.Reset()
	}
}

func (g *Game) Update() error { return nil }

func (g *Game) Layout(outerWindowWidth, outserWindowHeight int) (ScreenWidth, ScreenHeight int) {
	return outerWindowWidth, outserWindowHeight
}
```
