package mazebuilder

import (
	"github.com/harrytflv/go-design-patterns/creationaltemplate/maze"
)

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(room int)
	BuildDoor(rootFrom, roomTo int)
	GetMaze() maze.Maze
}
