package bombedmaze

import (
	"log"

	"github.com/harrytflv/go-design-patterns/abstractfactory/maze"
)

type bombedMazeFactory struct {
}

func (f *bombedMazeFactory) MakeMaze() maze.Maze {
	return &bombedMaze{}
}

func (f *bombedMazeFactory) MakeWall() maze.Wall {
	return &bombedWall{}
}

func (f *bombedMazeFactory) MakeRoom(n int) maze.Room {
	return &roomWithABomb{
		number: n,
		sides:  make(map[maze.Direction]maze.Wall),
	}
}

func (f *bombedMazeFactory) MakeDoor(r1, r2 maze.Room) maze.Door {
	return &door{
		from: r1,
		to:   r2,
	}
}

func (f *bombedMazeFactory) Created() {
	log.Println("A bombed maze is created.")
}

func NewBombedMazeFactory() maze.Factory {
	return &bombedMazeFactory{}
}
