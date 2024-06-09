package actorhandler

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
)

type IncomePacket struct {
	Opcode uint16
	Data   []byte
}

type Event struct {
	ActiveLogin string
}

func NewClient(conn net.Conn) actor.ActorHandle {
	ev := Event{
		ActiveLogin: "",
	}

	return func(me actor.ActorInterface, message any) any {
		incomePacket := message.(IncomePacket)
		ctx := context.Background()

		slog.Info("Income packet", "opcode", incomePacket.Opcode)

		var respPkt packets.PacketEncodeInterface

		switch packets.Opcode(incomePacket.Opcode) {
		case packets.OpcodeAuth:
			authPkt := &packets.Auth{}

			err := authPkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				return err
			}

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

			respPkt = cs

			slog.Info("Active login", "activeLogin", ev.ActiveLogin)
		case packets.OpcodeCreateCharacter:
			createCharPkt := packets.NewCharacterCreate()

			err := createCharPkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				return err
			}

			slog.Info("Create character", "name", createCharPkt.Name)

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

			respPkt = packets.NewCharacterCreateReply()
		case packets.OpcodeExit:
			conn.Close()
			return nil
		default:
			return fmt.Errorf("unknown opcode: %d", incomePacket.Opcode)
		}

		if respPkt == nil {
			return errors.New("response packet is nil")
		}

		pktBuf, err := packets.EncodeWithHeader(ctx, respPkt, binary.BigEndian)
		if err != nil {
			return err
		}

		_, err = conn.Write(pktBuf)
		if err != nil {
			return err
		}

		return nil
	}
}
