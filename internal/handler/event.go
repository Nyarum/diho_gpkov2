package handler

import (
	"bytes"
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

type Event struct {
	ActiveLogin string
}

func NewEventActor() actor.ActorHandle {
	fmt.Println("test2")
	ev := Event{
		ActiveLogin: "",
	}

	return func(me actor.ActorInterface, message any) any {
		incomePacket := message.(IncomePacket)
		ctx := context.Background()

		fmt.Println("Income packet:", incomePacket)

		switch packets.Opcode(incomePacket.Opcode) {
		case packets.OpcodeAuth:
			authPkt := &packets.Auth{}

			err := authPkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				fmt.Println("Error decoding auth packet:", err)
				return err
			}

			fmt.Println("Auth packet:", authPkt)

			ev.ActiveLogin = authPkt.Login

			actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), SaveAccount{
				Name: authPkt.Login,
				Data: *authPkt,
			})

			cs := packets.NewCharacterScreen()

			characters := actor.SendToMultipleAndGetOne(actor.ActorRegistry.GetByName("storage"), GetCharacters{
				Login: authPkt.Login,
			}).([]packets.Character)

			cs.CharacterLen = uint8(len(characters))
			cs.Characters = characters

			pktBuf, err := packets.EncodeWithHeader(ctx, cs, binary.BigEndian)
			if err != nil {
				fmt.Println("Error encoding cs packet:", err)
				return err
			}

			fmt.Println("Active login:", ev.ActiveLogin)

			incomePacket.Receiver.Send(sendToConn{
				buf: pktBuf,
			})
		case packets.OpcodeCreateCharacter:
			createCharPkt := packets.NewCharacterCreate()

			fmt.Println("Active login:", ev.ActiveLogin)

			err := createCharPkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				fmt.Println("Error decoding create character packet:", err)
				return err
			}

			fmt.Println("Create character packet:", createCharPkt)

			actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), SaveCharacter{
				Login: ev.ActiveLogin,
				Data: packets.Character{
					IsActive: true,
					Name:     createCharPkt.Name,
					Job:      "Newbie",
					Level:    1,
					LookSize: createCharPkt.LookSize,
					Look:     createCharPkt.Look,
				},
			})

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
