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
		if !g.AnyZeroes() {
			g.Reset()
		}

		_, err = ap.ReadOrResizeOrSignalOnce()
		if err != nil {
			log.FErrf("Error reading: %v", err)
		}
		slog.Info(fmt.Sprintf("%d\n", ap.Data[0]))
		switch ap.Data[0] {

		case 37, 'a': // left

			if g.Left() {
				g.AddOneInRandomSpot()
			}
		case 38, 'w': // up
			if g.Up() {
				g.AddOneInRandomSpot()
			}
		case 39, 'd': // right
			if g.Right() {
				g.AddOneInRandomSpot()
			}
		case 40, 's': // down
			if g.Down() {
				g.AddOneInRandomSpot()
			}
		case 'q':
			return
		}
	}

}
