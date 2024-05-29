package handler

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
)

func NewDataActor(conn net.Conn) actor.ActorHandle {
	return func(pid actor.PID, message any) any {
		defer conn.Close()

		buf := make([]byte, 2048)

		pktBuf, err := packets.EncodeWithHeader(packets.NewFirstTime(), binary.BigEndian)
		if err != nil {
			return err
		}

		_, err = conn.Write(pktBuf)
		if err != nil {
			return err
		}

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
