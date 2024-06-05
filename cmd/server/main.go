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

	handlerStorageActor, handlerStorageReturn := handler.NewStorageActor(ctx)
	if handlerStorageReturn.Err != nil {
		panic(handlerStorageReturn.Err)
	}
	defer handlerStorageReturn.DB.Close()

	actorStorage := actor.NewActor("storage", handlerStorageActor).Start(ctx)
	actorStorage.Send(actor.ActorReadyMessage)

	actor.ActorRegistry.Register(actorStorage)
	defer actor.ActorRegistry.Unregister(actorStorage)

	actor.NewActor("listener", handler.NewListenerActor(ctx, port)).Start(ctx).Send(actor.ActorReadyMessage)

	select {}
}
