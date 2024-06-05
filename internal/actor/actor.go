package actor

import (
	"context"

	"github.com/google/uuid"
)

type PID string
type ActorReady string

const (
	ActorNone PID = ""
)

var (
	ActorReadyMessage = ActorReady("ready")
)

type ActorHandle func(me ActorInterface, message any) any

type ActorInterface interface {
	Send(message any)
	SendReceive(message any) any
	Name() string
	PID() PID
}

type realtimeMessage struct {
	receiveCh chan any
	msg       any
}

type Actor struct {
	handle          ActorHandle
	mailbox         chan any
	realtimeMailbox chan realtimeMessage
	name            string
	cancel          chan struct{}
	pid             PID
}

func NewActor(name string, handle ActorHandle) Actor {
	return Actor{
		mailbox:         make(chan any, 1),
		realtimeMailbox: make(chan realtimeMessage, 1),
		name:            name,
		handle:          handle,
		cancel:          make(chan struct{}),
		pid:             PID(uuid.New().String()),
	}
}

func (a Actor) Name() string {
	return a.name
}

func (a Actor) PID() PID {
	return a.pid
}

func (a Actor) Send(message any) {
	a.mailbox <- message
}

func (a Actor) SendReceive(message any) any {
	realtimeMessage := realtimeMessage{
		receiveCh: make(chan any),
		msg:       message,
	}

	a.realtimeMailbox <- realtimeMessage
	return <-realtimeMessage.receiveCh
}

func (a Actor) Start(ctx context.Context) ActorInterface {
	go func() {
		for {
			select {
			case message := <-a.mailbox:
				a.handle(a, message)
			case realtimeMessage := <-a.realtimeMailbox:
				realtimeMessage.receiveCh <- a.handle(a, realtimeMessage.msg)
			case <-a.cancel:
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return a
}

func (a Actor) Stop(ctx context.Context) {
	a.cancel <- struct{}{}
}
