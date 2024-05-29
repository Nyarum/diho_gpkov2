package handler

import (
	"github.com/Nyarum/diho_gpkov2/internal/actor"
)

func NewEventActor() actor.ActorHandle {
	return func(pid actor.PID, message any) any {
		for {
			switch message.(type) {

			}
		}
	}
}
