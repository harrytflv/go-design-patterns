package enchantedmaze

import (
	"github.com/harrytflv/go-design-patterns/abstractfactory/maze"
)

type enchantedMaze struct {
	rooms []maze.Room
}

func (m *enchantedMaze) AddRoom(r maze.Room) {
	m.rooms = append(m.rooms, r)
}

type wall struct {
}

type enchantedRoom struct {
	number int
	sides  map[maze.Direction]maze.Wall
}

func (r *enchantedRoom) SetSide(d maze.Direction, w maze.Wall) {
	r.sides[d] = w
}

type doorNeedingSpell struct {
	from maze.Room
	to   maze.Room
}
