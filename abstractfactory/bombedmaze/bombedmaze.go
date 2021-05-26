package bombedmaze

import (
	"github.com/harrytflv/go-design-patterns/abstractfactory/maze"
)

type bombedMaze struct {
	rooms []maze.Room
}

func (m *bombedMaze) AddRoom(r maze.Room) {
	m.rooms = append(m.rooms, r)
}

type bombedWall struct {
}

type roomWithABomb struct {
	number int
	sides  map[maze.Direction]maze.Wall
}

func (r *roomWithABomb) SetSide(d maze.Direction, w maze.Wall) {
	r.sides[d] = w
}

type door struct {
	from maze.Room
	to   maze.Room
}
