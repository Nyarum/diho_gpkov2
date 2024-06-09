package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/actorhandler"
	"github.com/Nyarum/diho_gpkov2/internal/background"
)

func main() {
	port := ":1973"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	ctx := context.Background()

	handlerStorageActor, handlerStorageReturn := actorhandler.NewStorage(ctx)
	if handlerStorageReturn.Err != nil {
		panic(handlerStorageReturn.Err)
	}
	defer handlerStorageReturn.DB.Close()

	actorStorage := actor.NewActor("storage", handlerStorageActor).Start(ctx)
	actorStorage.Send(actor.ActorReadyMessage)

	actor.ActorRegistry.Register(actorStorage)
	defer actor.ActorRegistry.Unregister(actorStorage)

	err := background.NewTCP(ctx, port)
	if err != nil {
		slog.Error("Error", "error", err)
	}
}
