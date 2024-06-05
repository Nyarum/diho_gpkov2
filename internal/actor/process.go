package actor

func SendToMultiple(actors []ActorInterface, message any) {
	for _, actor := range actors {
		actor.Send(message)
	}
}

func SendToMultipleAndGetOne(actors []ActorInterface, message any) any {
	for _, actor := range actors {
		return actor.SendReceive(message)
	}

	return nil
}

func SendAndGetMultiple(actors []ActorInterface, message any) []any {
	results := make([]any, 0)

	for _, actor := range actors {
		results = append(results, actor.SendReceive(message))
	}

	return results
}
