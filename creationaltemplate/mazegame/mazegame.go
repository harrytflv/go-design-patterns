package mazegame

import (
	"github.com/harrytflv/go-design-patterns/creationaltemplate/direction"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/door"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/maze"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/room"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/wall"
)

type MazeGame struct {
}

func (g *MazeGame) CreateMaze() maze.Maze {
	aMaze := maze.New()
	r1 := room.New(1)
	r2 := room.New(1)
	theDoor := door.New(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(direction.North, wall.New())
	r1.SetSide(direction.East, theDoor)
	r1.SetSide(direction.South, wall.New())
	r1.SetSide(direction.West, wall.New())

	r2.SetSide(direction.North, wall.New())
	r2.SetSide(direction.East, wall.New())
	r2.SetSide(direction.South, wall.New())
	r2.SetSide(direction.West, theDoor)

	return aMaze
}
