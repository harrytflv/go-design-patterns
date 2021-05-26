package maze

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Maze interface {
	AddRoom(room Room)
}

type Wall interface {
}

type Room interface {
	SetSide(direction Direction, wall Wall)
}

type Door interface {
}
