package handler

import (
	"encoding/binary"
	"fmt"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
)

type IncomePacket struct {
	Opcode uint16
	Data   []byte
}

func NewEventActor() actor.ActorHandle {
	return func(pid actor.PID, message any) any {
		incomePacket := message.(IncomePacket)

		fmt.Println("Income packet:", incomePacket)

		switch incomePacket.Opcode {
		case 431:
			authPkt := &packets.Auth{}
			err := authPkt.Decode(incomePacket.Data, binary.BigEndian)
			if err != nil {
				fmt.Println("Error decoding auth packet:", err)
				return err
			}

			fmt.Println("Auth packet:", authPkt)
		}

		return nil
	}
}
