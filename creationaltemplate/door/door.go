package door

import (
	"github.com/harrytflv/go-design-patterns/creationaltemplate/room"
)

type Door interface {
}

type doorImpl struct {
	room1 room.Room
	room2 room.Room
}

func New(room1, room2 room.Room) Door {
	return doorImpl{
		room1: room1,
		room2: room2,
	}
}
