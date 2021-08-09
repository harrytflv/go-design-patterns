package mazegame

import (
	"github.com/harrytflv/go-design-patterns/builder/mazebuilder"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/maze"
)

type MazeGame interface {
	CreateMaze(builder mazebuilder.MazeBuilder) maze.Maze
	CreateComplexMaze(builder mazebuilder.MazeBuilder) maze.Maze
}

func New() MazeGame {
	return &mazeGameImpl{}
}

type mazeGameImpl struct {
}

func (g *mazeGameImpl) CreateMaze(builder mazebuilder.MazeBuilder) maze.Maze {
	builder.BuildMaze()
	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)
	return builder.GetMaze()
}

func (g mazeGameImpl) CreateComplexMaze(builder mazebuilder.MazeBuilder) maze.Maze {
	for i := 0; i < 1000; i++ {
		builder.BuildRoom(i)
	}
	return builder.GetMaze()
}
