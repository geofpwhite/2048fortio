package main

import (
	"fmt"
	"fortio/2048/game"

	"fortio.org/log"
	"fortio.org/terminal/ansipixels"
)

const (
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

var numColors = map[int]string{
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

func main() {
	ap := ansipixels.NewAnsiPixels(0)
	err := ap.Open()

	ap.WriteString(fmt.Sprintf("%d %d", ap.H, ap.W))
	if err != nil {
		log.FErrf("Error opening AnsiPixels: %v", err)
		panic("")
	}
	err = ap.GetSize()
	fmt.Println(ap.H, ap.W)
	if err != nil {
		fmt.Println(err)
		panic("")
	}

	defer ap.Restore()
	defer ap.ClearScreen()
	ap.ClearScreen()
	game := game.NewGame(ap)
	for {
		wi, hi := 0, 0

		for i := 0; i < ap.W-4; i += ap.W / 4 {
			hi = 0
			for j := 0; j < ap.H-4; j += ap.H / 4 {
				ap.StartSyncMode()
				ap.DrawColoredBox(i, j, ap.W/4, ap.H/4, numColors[game.State[wi][hi]], false)
				ap.WriteAtStr(i+ap.W/8, j+ap.H/8, fmt.Sprintf("%d", game.State[wi][hi]))
				// time.Sleep(200 * (time.Millisecond))
				ap.EndSyncMode()
				hi++
			}
			wi++
		}
		_, err = ap.ReadOrResizeOrSignalOnce()
		if err != nil {
			log.FErrf("Error reading: %v", err)
		}
		switch ap.Data[0] {
		case 37, 'a': // left
			game.Left()
			game.AddOneInRandomSpot()
		case 38, 'w': // up
			game.Up()
			game.AddOneInRandomSpot()
		case 39, 'd': // right
			game.Right()
			game.AddOneInRandomSpot()
		case 40, 's': // down
			game.Down()
			game.AddOneInRandomSpot()
		case 'q':
			return
		}
	}

}
