package handler

import (
	"context"
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

		_, eventActor := actor.NewActor("event", NewEventActor()).Start(context.Background())

		for {
			ln, err := conn.Read(buf)
			if err != nil {
				return err
			}

			if ln == 2 {
				conn.Write([]byte{0, 2})
				continue
			}

			fmt.Println("Has read from connection data:", string(buf[:ln]))
			fmt.Println("length:", ln)

			header, err := packets.DecodeHeader(buf)
			if err != nil {
				return err
			}

			buf = buf[8:]

			fmt.Println("Header:", header)

			if ln < int(header.Len) {
				moreData := make([]byte, int(header.Len)-ln)
				_, err := conn.Read(buf)
				if err != nil {
					return err
				}

				buf = append(buf, moreData...)
			}

			eventActor.Send(actor.ActorNone, IncomePacket{
				Opcode: header.Opcode,
				Data:   buf[:header.Len],
			})
		}
	}
}
