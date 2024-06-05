package handler

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
)

type dataMessage struct {
	header packets.Header
	buf    []byte
}

type sendToConn struct {
	buf []byte
}

type closeConn struct {
}

func NewDataActor(ctx context.Context, conn net.Conn) actor.ActorHandle {
	backgroundWork := func(me actor.ActorInterface) {
		buf := make([]byte, 2048)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				ln, err := conn.Read(buf)
				if err != nil {
					fmt.Println("Error reading from connection:", err)
					return
				}

				if ln == 2 {
					conn.Write([]byte{0, 2})
					continue
				}

				fmt.Println("Has read from connection data:", string(buf[:ln]))
				fmt.Println("length:", ln)

				header, err := packets.DecodeHeader(buf)
				if err != nil {
					fmt.Println("Error decoding header:", err)
					return
				}

				buf = buf[8:]

				fmt.Println("Header:", header)

				if ln < int(header.Len) {
					moreData := make([]byte, int(header.Len)-ln)
					_, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Error reading from connection more data:", err)
						return
					}

					buf = append(buf, moreData...)
				}

				me.Send(dataMessage{
					header: *header,
					buf:    buf[:header.Len],
				})
			}
		}
	}

	eventActor := actor.NewActor("event", NewEventActor()).Start(context.Background())

	return func(me actor.ActorInterface, message any) any {
		switch v := message.(type) {
		case dataMessage:
			eventActor.Send(IncomePacket{
				Receiver: me,
				Opcode:   v.header.Opcode,
				Data:     v.buf,
			})
		case sendToConn:
			fmt.Println("send to conn")
			conn.Write(v.buf)
		case closeConn:
			conn.Close()
		case actor.ActorReady:
			pktBuf, err := packets.EncodeWithHeader(ctx, packets.NewFirstTime(), binary.BigEndian)
			if err != nil {
				return err
			}

			_, err = conn.Write(pktBuf)
			if err != nil {
				return err
			}

			go backgroundWork(me)
		}

		return nil
	}
}
