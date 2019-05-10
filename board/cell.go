package board

type cell struct {
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


func (c *cell) addNeighbors(neighbors []*cell) {
	c.neighbors = append(c.neighbors, neighbors...)
}