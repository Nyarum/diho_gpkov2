package actor

func Send(registry *Registry, pid PID, message any) {
	registry.Get(pid).Send(message)
}

func SendReceive(registry *Registry, pid PID, message any) any {
	return registry.Get(pid).SendReceive(message)
}
