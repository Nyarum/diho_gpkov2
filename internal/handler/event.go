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
		ctx := context.Background()

		fmt.Println("Income packet:", incomePacket)

		switch packets.Opcode(incomePacket.Opcode) {
		case packets.OpcodeAuth:
			authPkt := &packets.Auth{}

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
		case packets.OpcodeCreateCharacter:
			createCharPkt := packets.NewCharacterCreate()

			err := createCharPkt.Decode(ctx, incomePacket.Data, binary.BigEndian)
			if err != nil {
				fmt.Println("Error decoding create character packet:", err)
				return err
			}

			fmt.Println("Create character packet:", createCharPkt)

			pktBuf, err := packets.EncodeWithHeader(ctx, packets.NewCharacterCreateReply(), binary.BigEndian)
			if err != nil {
				fmt.Println("Error encoding cs packet:", err)
				return err
			}

			incomePacket.Receiver.Send(sendToConn{
				buf: pktBuf,
			})
		case packets.OpcodeExit:
			incomePacket.Receiver.Send(closeConn{})
		default:
			fmt.Println("Unknown opcode:", incomePacket.Opcode)
		}

		return nil
	}
}
