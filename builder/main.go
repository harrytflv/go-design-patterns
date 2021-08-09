package main

import (
	"log"

	"github.com/harrytflv/go-design-patterns/builder/countingmazebuilder"
	"github.com/harrytflv/go-design-patterns/builder/mazegame"
	"github.com/harrytflv/go-design-patterns/builder/standardmazebuilder"
)

func main() {
	game := mazegame.New()
	builder := standardmazebuilder.New()

	game.CreateMaze(builder)
	maze := builder.GetMaze()

	maze.Go()

	countingBuilder := countingmazebuilder.New()
	game.CreateMaze(countingBuilder)
	log.Println(countingBuilder.GetCounts())
}
