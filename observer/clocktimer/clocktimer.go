package clocktimer

import (
	"context"
)

type ClockTimerObserver interface {
	Update(ctx context.Context, timer ClockTimer) error
}

type ClockTimer interface {
	Attach(o ClockTimerObserver) error
	Detach(o ClockTimerObserver) error
	// Whether Notify() is exposed envolves a discussion of "Who triggers the
	// update?" I choose calling notify internally instead of having clients
	// to do so for this particular case.
	// Notify(ctx context.Context) error

	GetSecond(ctx context.Context) (int, error)

	Tick(ctx context.Context) error
}

type clockTimerImpl struct {
	second int

	observers map[ClockTimerObserver]bool
}

func (t *clockTimerImpl) Attach(o ClockTimerObserver) error {
	t.observers[o] = true
	return nil
}

func (t *clockTimerImpl) Detach(o ClockTimerObserver) error {
	delete(t.observers, o)
	return nil
}

func (t *clockTimerImpl) GetSecond(ctx context.Context) (int, error) {
	return t.second, nil
}

func (t *clockTimerImpl) Tick(ctx context.Context) error {
	t.second += 1
	if err := t.notify(ctx); err != nil {
		return err
	}
	return nil
}

func (t *clockTimerImpl) notify(ctx context.Context) error {
	for o := range t.observers {
		if err := o.Update(ctx, t); err != nil {
			return err
		}
	}
	return nil
}

func NewClockTimer() ClockTimer {
	return &clockTimerImpl{
		observers: make(map[ClockTimerObserver]bool),
	}
}

var (
	_ ClockTimer = (*clockTimerImpl)(nil)
)
