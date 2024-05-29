package actor

import "sync"

type Registry struct {
	actors sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Get(pid PID) ActorInterface {
	if actor, ok := r.actors.Load(pid); ok {
		return actor.(ActorInterface)
	}
	return nil
}

func (r *Registry) Register(pid PID, actor ActorInterface) {
	r.actors.Store(pid, actor)
}

func (r *Registry) Unregister(pid PID) {
	r.actors.Delete(pid)
}
