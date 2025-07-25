package main

import (
	"fmt"
	"time"

	"fortio.org/log"
	"fortio.org/terminal/ansipixels"
)

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
	for i := 0; i < ap.W-4; i += ap.W / 4 {
		for j := 0; j < ap.H-4; j += ap.H / 4 {
			ap.ClearScreen()
			ap.StartSyncMode()
			ap.DrawSquareBox(i, j, ap.W/4, ap.H/4)
			ap.WriteAtStr(i+ap.W/8, j+ap.H/8, fmt.Sprintf("%d %d", i, j))
			time.Sleep(200 * (time.Millisecond))
			ap.EndSyncMode()
		}
	}

}
