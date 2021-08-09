package room

import (
	"sync"

	"github.com/harrytflv/go-design-patterns/creationaltemplate/direction"
	"github.com/harrytflv/go-design-patterns/creationaltemplate/mapsite"
)

type Room interface {
	mapsite.MapSite
	Number() int
	SetSide(d direction.Direction, s mapsite.MapSite)
	GetSide(d direction.Direction) mapsite.MapSite
}

func New(roomNumber int) Room {
	return &roomImpl{
		number: roomNumber,
	}
}

type roomImpl struct {
	number     int
	wallsMutex sync.RWMutex
	walls      map[direction.Direction]mapsite.MapSite
}

func (r *roomImpl) Number() int {
	return r.number
}

func (r *roomImpl) SetSide(d direction.Direction, s mapsite.MapSite) {
	r.wallsMutex.Lock()
	defer r.wallsMutex.Unlock()

	r.walls[d] = s
}

func (r *roomImpl) GetSide(d direction.Direction) mapsite.MapSite {
	r.wallsMutex.RLock()
	defer r.wallsMutex.RUnlock()

	return r.walls[d]
}
