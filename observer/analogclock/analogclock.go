package analogclock

import (
	"context"
	"log"
	"strings"

	"github.com/harrytflv/go-design-patterns/observer/clocktimer"
)

type AnalogClock interface {
	Draw(ctx context.Context) error

	Stop() error
}

type analogClockImpl struct {
	subject clocktimer.ClockTimer
}

func (c *analogClockImpl) Update(ctx context.Context, timer clocktimer.ClockTimer) error {
	if timer == c.subject {
		if err := c.Draw(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c *analogClockImpl) Draw(ctx context.Context) error {
	seconds, err := c.subject.GetSecond(ctx)
	if err != nil {
		return err
	}

	log.Println(strings.Repeat("|", seconds))

	return nil
}

func (c *analogClockImpl) Stop() error {
	return c.subject.Detach(c)
}

func NewAnalogClock(timer clocktimer.ClockTimer) (AnalogClock, error) {
	c := &analogClockImpl{
		subject: timer,
	}

	if err := c.subject.Attach(c); err != nil {
		return nil, err
	}

	return c, nil
}
