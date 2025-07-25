package main

import (
	"fmt"
	"log/slog"
	"os"

	"fortio.org/log"
	"fortio.org/terminal/ansipixels"
	"github.com/geofpwhite/2048fortio/game"
	// "github.com/geofpwhite/twenty48fortio/game"
)

func main() {
	file, err := os.Create("./event.log")
	// file, err := os.OpenFile("./event.log", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	loge := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{}))
	defer file.Close()
	slog.SetDefault(loge)
	ap := ansipixels.NewAnsiPixels(0)
	err = ap.Open()

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
	g := game.NewGame(ap)
	ap.HideCursor()
	g.Draw()
	for {
		if !g.AnyZeroes() && !g.AnyValidMoves() {
			g.Reset()
		}

		_, err = ap.ReadOrResizeOrSignalOnce()
		if err != nil {
			log.FErrf("Error reading: %v", err)
		}
		slog.Info(fmt.Sprintf("%d\n", ap.Data[0]))
		switch ap.Data[0] {
		case 27:
			switch ap.Data[2] {
			case 65:
				g.Up()
			case 66:
				g.Down()
			case 67:
				g.Right()
			case 68:
				g.Left()
			}
		case 37, 'a': // left
			g.Left()
		case 38, 'w': // up
			g.Up()
		case 39, 'd': // right
			g.Right()
		case 40, 's': // down
			g.Down()
		case 'q':
			return
		default:
			fmt.Println((ap.Data))
		}
	}

}
