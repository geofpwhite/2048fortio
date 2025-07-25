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
	time.Sleep(50 * time.Millisecond)
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
	// g.AP.ClearScreen()
	wi, hi := 0, 0
	for i := 0; i < g.AP.W-4; i += g.AP.W / 4 {
		hi = 0
		for j := 0; j < g.AP.H-4; j += g.AP.H / 4 {
			g.AP.StartSyncMode()
			g.AP.DrawRoundBox(0, 0, g.AP.W, g.AP.H)
			g.AP.DrawColoredBox(i+1, j+1, g.AP.W/4-2, g.AP.H/4-2, NumColors[g.State[wi][hi]], false)
			g.AP.WriteAtStr(i+g.AP.W/8, j+g.AP.H/8, "     ")
			g.AP.WriteAtStr(i+g.AP.W/8, j+g.AP.H/8, fmt.Sprintf("%d", g.State[wi][hi]))
			// time.Sleep(200 * (time.Millisecond))
			g.AP.EndSyncMode()
			hi++
		}
		wi++
	}
}

func (g *Game) Left() bool {
	//x--
	changed := false
	for y := 0; y < 4; y++ {
		for x := 1; x < 4; x++ {
			if g.State[x][y] == 0 {
				continue
			}
			var x2 int
			for x2 = x; x2 != 0 && g.State[x2][y] != 0 && g.State[x2-1][y] == 0; x2-- {
				g.State[x2][y], g.State[x2-1][y] = g.State[x2-1][y], g.State[x2][y]
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
			if x2 != 0 && g.State[x2][y] == g.State[x2-1][y] {
				g.State[x2][y], g.State[x2-1][y] = 0, g.State[x2-1][y]*2
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
		}
	}
	return changed
}

func (g *Game) Right() bool {
	changed := false
	for y := 0; y < 4; y++ {
		for x := 2; x > -1; x-- {
			if g.State[x][y] == 0 {
				continue
			}
			var x2 int
			for x2 = x; x2 != 3 && g.State[x2][y] != 0 && g.State[x2+1][y] == 0; x2++ {
				g.State[x2][y], g.State[x2+1][y] = g.State[x2+1][y], g.State[x2][y]
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
			if x2 != 3 && g.State[x2][y] == g.State[x2+1][y] {
				g.State[x2][y], g.State[x2+1][y] = 0, g.State[x2+1][y]*2
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
		}
	}
	return changed
}
func (g *Game) Up() bool {
	changed := false
	for x := 0; x < 4; x++ {
		for y := 1; y < 4; y++ {
			if g.State[x][y] == 0 {
				continue
			}
			var y2 int
			for y2 = y; y2 != 0 && g.State[x][y2] != 0 && g.State[x][y2-1] == 0; y2-- {
				g.State[x][y2], g.State[x][y2-1] = g.State[x][y2-1], g.State[x][y2]
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
			if y2 != 0 && g.State[x][y2] == g.State[x][y2-1] {
				g.State[x][y2], g.State[x][y2-1] = 0, g.State[x][y2-1]*2
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
		}
	}
	return changed

}
func (g *Game) Down() bool {
	changed := false
	for x := 0; x < 4; x++ {
		for y := 2; y > -1; y-- {
			if g.State[x][y] == 0 {
				continue
			}
			var y2 int
			for y2 = y; y2 != 3 && g.State[x][y2] != 0 && g.State[x][y2+1] == 0; y2++ {
				g.State[x][y2], g.State[x][y2+1] = g.State[x][y2+1], g.State[x][y2]
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
			if y2 != 3 && g.State[x][y2] == g.State[x][y2+1] {
				g.State[x][y2], g.State[x][y2+1] = 0, g.State[x][y2+1]*2
				g.Draw()
				time.Sleep(50 * time.Millisecond)
				changed = true
			}
		}
	}
	return changed
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

func (g *Game) Reset() {
	*g = *NewGame(g.AP)
}
