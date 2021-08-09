package countingmazebuilder

import (
	"sync/atomic"

	"github.com/harrytflv/go-design-patterns/builder/mazebuilder"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/maze"
)

type CountingMazeBuilder interface {
	mazebuilder.MazeBuilder
	GetCounts() (int, int)
}

type countingMazeBuilderI struct {
	doors       int32
	rooms       int32
	currentMaze maze.Maze
}

func New() CountingMazeBuilder {
	return &countingMazeBuilderI{}
}

func (mb *countingMazeBuilderI) BuildMaze() {
	mb.currentMaze = maze.New()
}

func (mb *countingMazeBuilderI) BuildRoom(roomNumber int) {
	atomic.AddInt32(&mb.rooms, 1)
}

func (mb *countingMazeBuilderI) BuildDoor(roomFrom, roomTo int) {
	atomic.AddInt32(&mb.doors, 1)
}

func (mb *countingMazeBuilderI) GetMaze() maze.Maze {
	return mb.currentMaze
}

func (mb *countingMazeBuilderI) GetCounts() (int, int) {
	return int(mb.doors), int(mb.rooms)
}
