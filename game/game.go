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
	State gameState
}

type gameState [4][4]int

func NewGame(ap *ansipixels.AnsiPixels) *Game {
	g := &Game{
		AP:    ap,
		State: gameState{},
	}
	g.AddOneInRandomSpot()
	return g
}

func (g *Game) AddOneInRandomSpot() {
	for {
		x, y := rand.IntN(4), rand.IntN(4)
		if g.State[x][y] != 0 {
			continue
		}
		g.State[x][y] = 1
		return
	}
}

func (g *Game) Left() {
	//x--
	for y := 0; y < 4; y++ {
		for x := 1; x < 4; x++ {
			var x2 int
			for x2 = x; x2 != 0 && g.State[x2][y] != 0 && g.State[x2-1][y] == 0; x2-- {
				g.State[x2][y], g.State[x2-1][y] = g.State[x2-1][y], g.State[x2][y]
			}
			if x2 != 0 && g.State[x2][y] == g.State[x2-1][y] {
				g.State[x2][y], g.State[x2-1][y] = 0, g.State[x2-1][y]*2
			}
		}
	}
}

func (g *Game) Right() {
	for y := 0; y < 4; y++ {
		for x := 2; x > -1; x-- {
			var x2 int
			for x2 = x; x2 != 3 && g.State[x2][y] != 0 && g.State[x2+1][y] == 0; x2++ {
				g.State[x2][y], g.State[x2+1][y] = g.State[x2+1][y], g.State[x2][y]
			}
			if x2 != 3 && g.State[x2][y] == g.State[x2+1][y] {
				g.State[x2][y], g.State[x2+1][y] = 0, g.State[x2+1][y]*2
			}
		}
	}
}
func (g *Game) Up() {
	for x := 0; x < 4; x++ {
		for y := 1; y < 4; y++ {
			var y2 int
			for y2 = y; y2 != 0 && g.State[x][y2] != 0 && g.State[x][y2-1] == 0; y2-- {
				g.State[x][y2], g.State[x][y2-1] = g.State[x][y2-1], g.State[x][y2]
			}
			if y2 != 0 && g.State[x][y2] == g.State[x][y2-1] {
				g.State[x][y2], g.State[x][y2-1] = 0, g.State[x][y2-1]*2
			}
		}
	}

}
func (g *Game) Down() {

	for x := 0; x < 4; x++ {
		for y := 2; y > -1; y-- {
			var y2 int
			for y2 = y; y2 != 3 && g.State[x][y2] != 0 && g.State[x][y2+1] == 0; y2++ {
				g.State[x][y2], g.State[x][y2+1] = g.State[x][y2+1], g.State[x][y2]
			}
			if y2 != 3 && g.State[x][y2] == g.State[x][y2+1] {
				g.State[x][y2], g.State[x][y2+1] = 0, g.State[x][y2+1]*2
			}
		}
	}

}
