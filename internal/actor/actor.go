package actor

import "github.com/google/uuid"

type PID string

type Actor struct {
}

func NewActor() Actor {
	return Actor{}
}

func (a Actor) Start() PID {
	return PID(uuid.New().String())
}

func (a Actor) Stop() {
}
