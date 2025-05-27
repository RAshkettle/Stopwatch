# Stopwatch

Go/Ebitengine Timer for games

Usage is pretty basic, just import into your application with
`go get github.com/RAshkettle/Stopwatch`

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
