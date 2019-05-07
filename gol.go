package main

import (
	"awesomeProject/board"
	"fmt"
	"image"
	"image/gif"
	"os"
	"runtime"
)

func main() {
	field := board.NewField(500, 500, 3)
	var images []*image.Paletted
	var delays []int
	var mem runtime.MemStats
	frames := 1000
	for i := 0; i < frames; i++ {
		field.Update()
		img := field.GetCurrentImage()

		images = append(images, img)
		delays = append(delays, 0)
		if i % (frames / 100) == 0 {
			fmt.Printf("%d%%", (i/(frames/100)) + 1)
			runtime.ReadMemStats(&mem)
			fmt.Printf(" Total Allocated Memory = %v MiB\n", bToMb(mem.TotalAlloc))
		}
	}
	f, _ := os.OpenFile("game_of_life.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
