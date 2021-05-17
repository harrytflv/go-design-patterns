package main

import (
	"context"
	"log"
	"time"

	"github.com/harrytflv/go-design-patterns/observer/analogclock"
	"github.com/harrytflv/go-design-patterns/observer/clocktimer"
	"github.com/harrytflv/go-design-patterns/observer/digitalclock"
)

func main() {
	timer := clocktimer.NewClockTimer()
	analogClock, err := analogclock.NewAnalogClock(timer)
	if err != nil {
		panic(err)
	}
	digitalClock, err := digitalclock.NewDigitalClock(timer)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		log.Println("Timer Ticked.")
		if err := timer.Tick(context.Background()); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}

	log.Println("Let's see the two clocks once again.")
	if err := analogClock.Draw(context.Background()); err != nil {
		panic(err)
	}
	if err := digitalClock.Draw(context.Background()); err != nil {
		panic(err)
	}
}
