package actor

import (
	"context"
	"log/slog"
	"runtime/debug"

	"github.com/google/uuid"
)

type PID string
type ActorReady string
type ActorInternalState int

const (
	ActorNone     PID                = ""
	ActorRestored ActorInternalState = 1
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
	var actor func()

	actor = func() {
		handler := func() {
			for {
				select {
				case message := <-a.mailbox:
					if v := a.handle(a, message); v != nil {
						errType, ok := v.(error)
						if ok {
							slog.Error("Can't handle message", "error", errType, "pid", a.pid, "message", message)
						}
					}
				case realtimeMessage := <-a.realtimeMailbox:
					realtimeMessage.receiveCh <- a.handle(a, realtimeMessage.msg)
				case <-a.cancel:
					return
				case <-ctx.Done():
					return
				}
			}
		}

		recover := func() {
			if r := recover(); r != nil {
				slog.Warn("Supervisor recovered", "error", r, "stack", string(debug.Stack()))
				a.mailbox <- ActorRestored
				actor()
			}
		}

		defer recover()
		handler()
	}

	go actor()

	return a
}

func (a Actor) Stop(ctx context.Context) {
	a.cancel <- struct{}{}
}
