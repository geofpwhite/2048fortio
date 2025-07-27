package game

import (
	"fmt"
	"math/rand/v2"
	"time"

	"fortio.org/terminal/ansipixels"
)

type input uint

const (
	ZERO                  = ansipixels.Black
	ONE                   = ansipixels.Red
	TWO                   = ansipixels.Green
	FOUR                  = ansipixels.Yellow
	EIGHT                 = ansipixels.Blue
	SIXTEEN               = ansipixels.Purple
	THIRTYTWO             = ansipixels.Cyan
	SIXTYFOUR             = ansipixels.Gray
	ONEHUNDREDTWENTYEIGHT = ansipixels.DarkGray
	TWOHUNDREDFIFTYSIX    = ansipixels.BrightGreen
)

var NumColors = map[int]string{
	0:   ZERO,
	1:   ONE,
	2:   TWO,
	4:   FOUR,
	8:   EIGHT,
	16:  SIXTEEN,
	32:  THIRTYTWO,
	64:  SIXTYFOUR,
	128: ONEHUNDREDTWENTYEIGHT,
	256: TWOHUNDREDFIFTYSIX,
}

const (
	UP input = iota
	DOWN
	LEFT
	RIGHT
)

type Game struct {
	AP           *ansipixels.AnsiPixels
	State        gameState
	ShowControls bool
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

func (g *Game) Score() int {
	var ret int
	for _, row := range g.State {
		for _, num := range row {
			ret += num
		}
	}
	return ret
}

func (g *Game) AddOneInRandomSpot() {
	time.Sleep(15 * time.Millisecond)
	for {
		x, y := rand.IntN(4), rand.IntN(4)
		if g.State[x][y] != 0 {
			continue
		}
		g.State[x][y] = 1
		g.Draw()
		return
	}
}
func (g *Game) Draw() {
	wi, hi := 0, 0
	for i := 0; i < g.AP.W-4; i += g.AP.W / 4 {
		hi = 0
		for j := 3; j < g.AP.H-4; j += g.AP.H / 4 {
			g.AP.StartSyncMode()
			g.AP.DrawColoredBox(i+1, j+1, g.AP.W/4-2, g.AP.H/4-2, NumColors[g.State[wi][hi]], false)
			g.AP.WriteAtStr(i+g.AP.W/8, j+g.AP.H/8, "     ")
			g.AP.WriteAtStr(i+g.AP.W/8, j+g.AP.H/8, fmt.Sprintf("%d", g.State[wi][hi]))
			// time.Sleep(200 * (time.Millisecond))
			g.AP.EndSyncMode()
			hi++
		}
		wi++
	}
	if g.ShowControls {
		g.AP.ClearScreen()
		g.AP.StartSyncMode()
		g.AP.ClearScreen()
		g.AP.DrawRoundBox(g.AP.W/4, g.AP.H/4, g.AP.W/2, g.AP.H/2)
		g.AP.WriteAtStr(g.AP.W/3, g.AP.H/2, "press wasd or arrow keys to move")
		g.AP.EndSyncMode()
		return
	}
	g.AP.StartSyncMode()
	g.AP.DrawRoundBox(0, 0, 8, 3)
	g.AP.WriteAtStr(1, 1, fmt.Sprintf("%s%d", ansipixels.Green, g.Score()))
	g.AP.EndSyncMode()
}

func (g *Game) shift(x0, y0, xf, yf, xne, yne, dx, dy, dx1, dy1 int) bool {
	changed := false
	for x := x0; x != xf; x += dx1 {
		for y := y0; y != yf; y += dy1 {
			if g.State[x][y] == 0 {
				continue
			}
			var y2 int
			for y2 = y; y2 != yne && g.State[x][y2] != 0 && g.State[x][y2+dy] == 0 && dy != 0; y2 += dy {
				g.State[x][y2], g.State[x][y2+dy] = g.State[x][y2+dy], g.State[x][y2]
				g.Draw()
				time.Sleep(15 * time.Millisecond)
				changed = true
			}
			if y2 != yne && g.State[x][y2] == g.State[x][y2+dy] && dy != 0 {
				g.State[x][y2], g.State[x][y2+dy] = 0, g.State[x][y2+dy]*2
				g.Draw()
				time.Sleep(15 * time.Millisecond)
				changed = true
			}
			var x2 int
			for x2 = x; x2 != xne && g.State[x2][y] != 0 && g.State[x2+dx][y] == 0 && dx != 0; x2 += dx {
				g.State[x2][y], g.State[x2+dx][y] = g.State[x2+dx][y], g.State[x2][y]
				g.Draw()
				time.Sleep(15 * time.Millisecond)
				changed = true
			}
			if x2 != xne && g.State[x2][y] == g.State[x2+dx][y] && dx != 0 {
				g.State[x2][y], g.State[x2+dx][y] = 0, g.State[x2+dx][y]*2
				g.Draw()
				time.Sleep(15 * time.Millisecond)
				changed = true
			}
		}
	}
	return changed

}

func (g *Game) Left() {
	if g.shift(1, 0, 4, 4, 0, 3, -1, 0, 1, 1) {
		g.AddOneInRandomSpot()
	}
}
func (g *Game) Right() {
	if g.shift(2, 0, -1, 4, 3, 0, 1, 0, -1, 1) {
		g.AddOneInRandomSpot()
	}
}

func (g *Game) Up() {
	if g.shift(0, 1, 4, 4, 0, 0, 0, -1, 1, 1) {
		g.AddOneInRandomSpot()
	}
}
func (g *Game) Down() {
	if g.shift(0, 2, 4, -1, 0, 3, 0, 1, 1, -1) {
		g.AddOneInRandomSpot()
	}
}

func (g *Game) AnyZeroes() bool {
	for _, row := range g.State {
		for _, cell := range row {
			if cell == 0 {
				return true
			}
		}
	}
	return false
}

func (g *Game) AnyValidMoves() bool {
	// search down and right
	for x := range 3 {
		for y := range 3 {
			if g.State[x][y] == g.State[x+1][y] || g.State[x][y+1] == g.State[x][y] {
				return true
			}
		}
	}
	return false
}

func (g *Game) Reset() {
	*g = *NewGame(g.AP)
}
