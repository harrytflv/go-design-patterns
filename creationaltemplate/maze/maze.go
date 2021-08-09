package maze

import (
	"log"
	"sync"

	"github.com/harrytflv/go-design-patterns/creationaltemplate/room"
)

type Maze interface {
	Go()
	AddRoom(room room.Room)
	RoomNo(roomNumber int) room.Room
}

func New() Maze {
	return &mazeImpl{
		rooms: make(map[int]room.Room),
	}
}

type mazeImpl struct {
	roomMutex sync.RWMutex
	rooms     map[int]room.Room
}

func (m *mazeImpl) Go() {
	log.Println("Go!")
}

func (m *mazeImpl) AddRoom(room room.Room) {
	m.roomMutex.Lock()
	defer m.roomMutex.Unlock()

	m.rooms[room.Number()] = room
}

func (m *mazeImpl) RoomNo(roomNumber int) room.Room {
	m.roomMutex.RLock()
	defer m.roomMutex.RUnlock()

	return m.rooms[roomNumber]
}
