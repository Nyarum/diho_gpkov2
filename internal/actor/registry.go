package actor

import (
	"sync"
)

var (
	ActorRegistry = NewRegistry()
)

type Registry struct {
	actorsByName sync.Map
	actorsByPID  sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) GetByPID(pid PID) ActorInterface {
	if actor, ok := r.actorsByPID.Load(pid); ok {
		return actor.(ActorInterface)
	}
	return nil
}

func (r *Registry) GetByName(name string) []ActorInterface {
	if actor, ok := r.actorsByName.Load(name); ok {
		if actors, ok := actor.([]ActorInterface); ok {
			return actors
		}
	}
	return nil
}

func (r *Registry) Register(actor ActorInterface) {
	r.actorsByPID.Store(actor.PID(), actor)

	if actors, ok := r.actorsByName.LoadOrStore(actor.Name(), []ActorInterface{actor}); ok {
		if actorsSlice, ok := actors.([]ActorInterface); ok {
			actorsSlice = append(actorsSlice, actor)
			r.actorsByName.Store(actor.Name(), actorsSlice)
		}
	}
}

func (r *Registry) Unregister(actor ActorInterface) {
	r.actorsByPID.Delete(actor.PID())
	r.actorsByName.Delete(actor.Name())
}
