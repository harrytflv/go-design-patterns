package mazegame

import (
	"github.com/harrytflv/go-design-patterns/abstractfactory/maze"
)

func CreateMaze(factory maze.Factory) maze.Maze {
	aMaze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	aDoor := factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(maze.North, factory.MakeWall())
	r1.SetSide(maze.East, aDoor)
	r1.SetSide(maze.South, factory.MakeWall())
	r1.SetSide(maze.West, factory.MakeWall())

	r2.SetSide(maze.North, factory.MakeWall())
	r2.SetSide(maze.East, factory.MakeWall())
	r2.SetSide(maze.South, factory.MakeWall())
	r2.SetSide(maze.West, aDoor)

	factory.Created()

	return aMaze
}
