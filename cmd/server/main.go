package main

import (
	"context"
	"os"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/handler"
)

func main() {
	port := ":1973"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	ctx := context.Background()

	_, listenerActor := actor.NewActor("listener", handler.NewListenerActor(ctx, port)).Start(ctx)
	listenerActor.Send(actor.ActorNone, "ready")

	select {}
}
