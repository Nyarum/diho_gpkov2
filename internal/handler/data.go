package handler

import (
	"fmt"
	"net"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
)

func NewDataActor(conn net.Conn) actor.ActorHandle {
	return func(pid actor.PID, message any) any {
		defer conn.Close()

		buf := make([]byte, 2048)

		for {
			ln, err := conn.Read(buf)
			if err != nil {
				return err
			}

			fmt.Println("Has read from connection data:", string(buf[:ln]))
			fmt.Println("length:", ln)

		}
	}
}
