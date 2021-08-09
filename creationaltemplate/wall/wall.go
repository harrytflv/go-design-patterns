package wall

import (
	"github.com/harrytflv/go-design-patterns/creationaltemplate/mapsite"
)

type Wall interface {
	mapsite.MapSite
}

type wallImpl struct {
}

func New() Wall {
	return &wallImpl{}
}
