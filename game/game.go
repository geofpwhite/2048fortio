package game

import (
	"math/rand/v2"

	"fortio.org/terminal/ansipixels"
)

type input uint

const (
	UP input = iota
	DOWN
	LEFT
	RIGHT
)

type Game struct {
	AP    *ansipixels.AnsiPixels
	state gameState
}

type gameState [4][4]int

func (g *Game) AddOneInRandomSpot() {
	for {
		x, y := rand.IntN(4), rand.IntN(4)
		if g.state[x][y] != 0 {
			continue
		}
		g.state[x][y] = 1
		return
	}
}

func (g *Game) Left() {

}
