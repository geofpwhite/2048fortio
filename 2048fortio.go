package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"fortio.org/log"
	"fortio.org/terminal/ansipixels"
	"github.com/geofpwhite/2048fortio/game"
	// "github.com/geofpwhite/twenty48fortio/game"
)

func main() {
	usrhomedir, err := os.UserHomeDir()
	if err != nil {
		panic("")
	}
	hs := 0
	hsFile, err := os.ReadFile(usrhomedir + "/highscore.txt")

	if err == nil {
		str := string(hsFile)
		num, err := strconv.Atoi(strings.Trim(str, "\n\r"))
		hs = num
		if err != nil {
			hs = 0
		}
	}
	file, err := os.Create(usrhomedir + "/event.log")
	if err != nil {
		fmt.Println(err)
		panic("")
	}
	loge := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{}))
	defer file.Close()
	slog.SetDefault(loge)
	ap := ansipixels.NewAnsiPixels(0)

	err = ap.Open()
	ap.SharedInput.Start(context.TODO())

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
	g := game.NewGame(ap, hs)
	defer func() {
		if g.HighScore > hs {
			file, err := os.Create(usrhomedir + "/highscore.txt")
			if err != nil {
				return
			}
			file.Write([]byte(fmt.Sprintf("%d", g.HighScore)))
			file.Close()
		}
	}()
	ap.OnResize = func() error {
		slog.Info(fmt.Sprintf("resized h=%d w=%d", ap.H, ap.W))
		newGame := game.NewGame(ap, hs)
		newGame.State = g.State
		g = newGame
		return nil
	}
	ap.HideCursor()
	g.Draw()
	for {
		h, w := ap.H, ap.W
		err = ap.GetSize()
		if err != nil {
			fmt.Println(err)
			panic("")
		}
		if w != ap.W || h != ap.H {
			ap.ClearScreen()
		}
		if !g.AnyZeroes() && !g.AnyValidMoves() {
			g.Reset()
		}

		_, err = ap.ReadOrResizeOrSignalOnce()
		if err != nil {
			log.FErrf("Error reading: %v", err)
		}
		if len(ap.Data) == 0 {
			continue
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
		case 'h', '?':
			g.ShowControls = !g.ShowControls
			ap.ClearScreen()
			g.Draw()
		case 'q':
			return

		}
	}

}
