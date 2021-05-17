package digitalclock

import (
	"context"
	"log"

	"github.com/harrytflv/go-design-patterns/observer/clocktimer"
)

type DigitalClock interface {
	clocktimer.ClockTimerObserver

	Draw(ctx context.Context) error

	Stop() error
}

type digitalClockImpl struct {
	subject clocktimer.ClockTimer
}

func (c *digitalClockImpl) Update(ctx context.Context, timer clocktimer.ClockTimer) error {
	if timer == c.subject {
		if err := c.Draw(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (c *digitalClockImpl) Draw(ctx context.Context) error {
	seconds, err := c.subject.GetSecond(ctx)
	if err != nil {
		return err
	}

	log.Printf("Second: %d", seconds)

	return nil
}

func (c *digitalClockImpl) Stop() error {
	return c.subject.Detach(c)
}

func NewDigitalClock(timer clocktimer.ClockTimer) (DigitalClock, error) {
	c := &digitalClockImpl{
		subject: timer,
	}

	if err := c.subject.Attach(c); err != nil {
		return nil, err
	}

	return c, nil
}
