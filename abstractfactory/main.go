package main

import (
	"github.com/harrytflv/go-design-patterns/abstractfactory/bombedmaze"
	"github.com/harrytflv/go-design-patterns/abstractfactory/enchantedmaze"
	"github.com/harrytflv/go-design-patterns/abstractfactory/mazegame"
)

func main() {
	mazegame.CreateMaze(bombedmaze.NewBombedMazeFactory())
	mazegame.CreateMaze(enchantedmaze.NewEnchantedMazeFactory())
}
