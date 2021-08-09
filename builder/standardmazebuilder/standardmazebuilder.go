package standardmazebuilder

import (
	"github.com/harrytflv/go-design-patterns/builder/mazebuilder"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/direction"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/door"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/maze"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/room"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/wall"
)

func New() mazebuilder.MazeBuilder {
	return &standardMazeBuilder{}
}

type standardMazeBuilder struct {
	currentMaze maze.Maze
}

func (mb *standardMazeBuilder) BuildMaze() {
	mb.currentMaze = maze.New()
}

func (mb *standardMazeBuilder) BuildRoom(roomNumber int) {
	if mb.currentMaze.RoomNo(roomNumber) != nil {
		return
	}
	r := room.New(roomNumber)
	mb.currentMaze.AddRoom(r)

	r.SetSide(direction.North, wall.New())
	r.SetSide(direction.South, wall.New())
	r.SetSide(direction.East, wall.New())
	r.SetSide(direction.West, wall.New())
}

func (mb *standardMazeBuilder) BuildDoor(roomFrom, roomTo int) {
	r1 := mb.currentMaze.RoomNo(roomFrom)
	r2 := mb.currentMaze.RoomNo(roomTo)
	d := door.New(r1, r2)

	r1.SetSide(mb.commonWall(r1, r2), d)
	r2.SetSide(mb.commonWall(r1, r2), d)
}

func (mb *standardMazeBuilder) GetMaze() maze.Maze {
	return mb.currentMaze
}

func (mb *standardMazeBuilder) commonWall(room1, room2 room.Room) direction.Direction {
	return direction.North
}
