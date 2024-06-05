package handler

import (
	"context"
	"fmt"
	"net"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
)

func NewListenerActor(ctx context.Context, port string) actor.ActorHandle {
	backgroundWork := func(listener net.Listener) error {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					return err
				}

				fmt.Println("accepted connect")

				actor.NewActor("data", NewDataActor(ctx, conn)).Start(ctx).Send(actor.ActorReadyMessage)
			}
		}
	}

	return func(me actor.ActorInterface, message any) any {
		switch message.(type) {
		case actor.ActorReady:
			listener, err := net.Listen("tcp", port)
			if err != nil {
				fmt.Println("Error creating listener:", err)
				return err
			}

			defer listener.Close()
			fmt.Printf("Server is listening on port %s\n", port)

			return backgroundWork(listener)
		}

		return nil
	}
}
