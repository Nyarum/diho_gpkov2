package actor

import (
	"context"

	"github.com/google/uuid"
)

type PID string

const (
	ActorNone PID = ""
)

type ActorHandle func(pid PID, message any) any

type ActorInterface interface {
	Send(pid PID, message any)
	SendReceive(pid PID, message any) any
	Name() string
}

type Actor struct {
	handle  ActorHandle
	mailbox chan any
	name    string
	cancel  chan struct{}
}

func NewActor(name string, handle ActorHandle) Actor {
	return Actor{
		mailbox: make(chan any, 1),
		name:    name,
		handle:  handle,
		cancel:  make(chan struct{}),
	}
}

func (a Actor) Name() string {
	return a.name
}

func (a Actor) Send(pid PID, message any) {
	a.mailbox <- message
}

func (a Actor) SendReceive(pid PID, message any) any {
	return a.handle(pid, message)
}

func (a Actor) Start(ctx context.Context) (PID, ActorInterface) {
	pid := PID(uuid.New().String())

	go func() {
		for {
			select {
			case message := <-a.mailbox:
				a.handle(pid, message)
			case <-a.cancel:
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return pid, a
}

func (a Actor) Stop(ctx context.Context) {
	a.cancel <- struct{}{}
}
