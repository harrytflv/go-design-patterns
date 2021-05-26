package maze

type Factory interface {
	MakeMaze() Maze
	MakeWall() Wall
	MakeRoom(n int) Room
	MakeDoor(r1, r2 Room) Door

	Created()
}
