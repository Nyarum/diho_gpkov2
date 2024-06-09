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
		switch v := message.(type) {
		case actor.ActorInternalState:
			if v == actor.ActorRestored {
				conn.Close()
				slog.Info("I'm restored")
			}

			return nil
		}

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

			resp := actor.SendToMultipleAndGetOne(actor.ActorRegistry.GetByName("storage"), GetAccount{
				Name: authPkt.Login,
			})

			var account packets.Auth
			if resp == nil {
				saveAuthPkt := *authPkt

				actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), SaveAccount{
					Name: authPkt.Login,
					Data: saveAuthPkt,
				})

				account = *authPkt
			} else {
				account = resp.(packets.Auth)
			}

			ev.ActiveLogin = authPkt.Login

			cs := packets.NewCharacterScreen()

			characters := actor.SendToMultipleAndGetOne(actor.ActorRegistry.GetByName("storage"), GetCharacters{
				Login: authPkt.Login,
			}).([]packets.Character)

			cs.CharacterLen = uint8(len(characters))
			cs.Characters = characters

			if len(account.PincodeHash) != 0 {
				cs.Pincode = 1
			}

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
		case packets.OpcodeRemoveCharacter:
			pkt := packets.NewCharacterRemove()

			err := pkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				return err
			}

			slog.Info("Remove character", "pkt", pkt)

			actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), RemoveCharacter{
				Login: ev.ActiveLogin,
				Name:  pkt.Name,
			})

			respPkt = packets.NewCharacterRemoveReply()
		case packets.OpcodeCreatePincode:
			pkt := packets.NewCreatePincode()

			err := pkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				return err
			}

			slog.Info("Create pincode", "pkt", pkt)

			actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), UpdatePincode{
				Login: ev.ActiveLogin,
				Hash:  pkt.Hash,
			})

			respPkt = packets.NewCreatePincodeReply()
		case packets.OpcodeChangePincode:
			pkt := packets.NewUpdatePincode()

			err := pkt.Decode(ctx, bytes.NewReader(incomePacket.Data), binary.BigEndian)
			if err != nil {
				return err
			}

			slog.Info("Update pincode", "pkt", pkt)

			actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), UpdatePincode{
				Login: ev.ActiveLogin,
				Hash:  pkt.Hash,
			})

			respPkt = packets.NewUpdatePincodeReply()
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
