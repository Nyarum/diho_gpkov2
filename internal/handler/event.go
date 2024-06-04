package handler

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
)

type IncomePacket struct {
	Receiver actor.ActorInterface
	Opcode   uint16
	Data     []byte
}

func NewEventActor() actor.ActorHandle {
	return func(me actor.ActorInterface, message any) any {
		incomePacket := message.(IncomePacket)

		fmt.Println("Income packet:", incomePacket)

		switch incomePacket.Opcode {
		case 431:
			authPkt := &packets.Auth{}
			ctx := context.Background()

			err := authPkt.Decode(ctx, incomePacket.Data, binary.BigEndian)
			if err != nil {
				fmt.Println("Error decoding auth packet:", err)
				return err
			}

			fmt.Println("Auth packet:", authPkt)

			pktBuf, err := packets.EncodeWithHeader(ctx, packets.NewCharacterScreen(), binary.BigEndian)
			if err != nil {
				fmt.Println("Error encoding cs packet:", err)
				return err
			}

			incomePacket.Receiver.Send(sendToConn{
				buf: pktBuf,
			})
		case 432:
			incomePacket.Receiver.Send(closeConn{})
		}

		return nil
	}
}
