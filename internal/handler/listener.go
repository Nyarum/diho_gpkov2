package handler

import (
	"context"
	"fmt"
	"net"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
)

func NewListenerActor(ctx context.Context, port string) actor.ActorHandle {
	return func(pid actor.PID, message any) any {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			fmt.Println("Error creating listener:", err)
			return err
		}

		defer listener.Close()
		fmt.Printf("Server is listening on port %s\n", port)

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				return err
			}

			actor.NewActor("data", NewDataActor(conn)).Start(ctx)
		}
	}
}
