package board

import (
	"fmt"
)

type cell struct {
	col, row int
	alive bool
	neighbors []*cell
}

func (c cell) next() bool {
	alive := 0
	for _, neighbor := range c.neighbors {
		if neighbor.alive {
			alive++
		}
	}
	return (c.alive && alive == 2)|| alive == 3
}

func (c *cell) update(status bool) {
	c.alive = status
}

func (c *cell) addNeighbor(neighbor *cell) {
	c.neighbors = append(c.neighbors, neighbor)
}

func (c *cell) addNeighbors(neighbors []*cell) {
	c.neighbors = append(c.neighbors, neighbors...)
}

func (c cell) printNeighbors(currentState bool, nextState bool) {
	str := ""
	for i, neighbor := range c.neighbors {
		if neighbor.alive {
			str += "1 "
		} else {
			str += "0 "
		}

		switch i {
			case 2:
				str += "\n"
				break
			case 3:
				if currentState {
					str += "1 "
				} else {
					str += "0 "
				}
				break
			case 4:
				str += "\n"
		}
	}
	fmt.Printf("Current: %t, Next: %t, Location: [%d,%d]:\n", currentState, nextState, c.col, c.row)
	fmt.Println(str + "\n")
}