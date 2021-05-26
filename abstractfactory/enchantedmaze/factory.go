package enchantedmaze

import (
	"log"

	"github.com/harrytflv/go-design-patterns/abstractfactory/maze"
)

type enchantedMazeFactory struct {
}

func (f *enchantedMazeFactory) MakeMaze() maze.Maze {
	return &enchantedMaze{}
}

func (f *enchantedMazeFactory) MakeWall() maze.Wall {
	return &wall{}
}

func (f *enchantedMazeFactory) MakeRoom(n int) maze.Room {
	return &enchantedRoom{
		number: n,
		sides:  make(map[maze.Direction]maze.Wall),
	}
}

func (f *enchantedMazeFactory) MakeDoor(r1, r2 maze.Room) maze.Door {
	return &doorNeedingSpell{
		from: r1,
		to:   r2,
	}
}

func (f *enchantedMazeFactory) Created() {
	log.Println("An enchanted maze is created.")
}

func NewEnchantedMazeFactory() maze.Factory {
	return &enchantedMazeFactory{}
}
