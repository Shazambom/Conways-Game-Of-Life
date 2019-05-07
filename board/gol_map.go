package board

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"math/rand"
	"strings"
	"time"
)

type field struct {
	height, width int
	nextState [][]bool
	cells [][]*cell
	scale int
}

func NewField(height int, width int, scale int) *field {
	f := field{height, width,make([][]bool, height), make([][]*cell, height), scale}
	for i := range f.nextState {
		f.nextState[i] = make([]bool, width)
		f.cells[i] = make([]*cell, width)
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			alive := random.Int31()&(1<<30) == 0
			f.nextState[i][j] = alive
			f.cells[i][j] = &cell{ i, j,alive, make([]*cell, 0)}
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// 1 2 3
			// 4 # 5
			// 6 7 8
			neighbors := []*cell{
				f.getSafeCell(i - 1, j - 1),
				f.getSafeCell(i, j - 1),
				f.getSafeCell(i + 1, j - 1),
				f.getSafeCell(i - 1, j),
				f.getSafeCell(i + 1, j),
				f.getSafeCell(i - 1, j + 1),
				f.getSafeCell(i, j + 1),
				f.getSafeCell(i + 1, j + 1)}
			f.cells[i][j].addNeighbors(neighbors)
		}
	}

	return &f
}

func (f *field) getSafeCell(col int, row int) *cell {
	r := row + f.width
	r %= f.width
	c := col + f.height
	c %= f.height
	return f.cells[c][r]
}


func (f *field) Update() {
	for r, row := range f.cells {
		for c, cell := range row {
			f.nextState[r][c] = cell.next()
		}
	}

	for r, row := range f.cells {
		for c, cell := range row {
			cell.update(f.nextState[r][c])
		}
	}
}

func (f *field) PrintField() {
	intArr := make([][]uint8, f.height)
	for i := 0; i < f.height; i++ {
		intArr[i] = make([]uint8, f.width)
	}
	str := ""
	for i, row := range f.cells {
		for j, val := range row {
			intArr[i][j] = btou(val.alive)
		}
	}
	for _, row := range intArr {
		str += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(row)), ""), "[]")
		str += "\n"
	}
	fmt.Println(str + "\n\n")
}

func btou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func (f *field) GetCurrentImage() *image.Paletted {
	img := image.NewPaletted(image.Rectangle{image.Point{0,0}, image.Point{f.width * f.scale, f.height * f.scale}}, palette.Plan9)
	for i := 0; i < f.height; i++ {
		for j := 0; j < f.width; j++ {
			if f.cells[i][j].alive {
				for k := 0; k < f.scale; k++ {
					for l := 0; l < f.scale; l++ {
						img.Set((j * f.scale) + l, (i * f.scale) + k, color.RGBA{255, 255, 0, 0xff})
					}
				}
			} else {
				for k := 0; k < f.scale; k++ {
					for l := 0; l < f.scale; l++ {
						img.Set((j * f.scale) + l, (i * f.scale) + k, color.Black)
					}
				}
			}

		}
	}
	return img
}
