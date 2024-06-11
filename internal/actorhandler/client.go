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

type TempData struct {
	ActiveLogin string
}

func NewClient(conn net.Conn) actor.ActorHandle {
	tempData := TempData{
		ActiveLogin: "",
	}

	return func(me actor.ActorInterface, message any) any {
		logger := slog.With("pid", me.PID())

		switch v := message.(type) {
		case actor.ActorInternalState:
			if v == actor.ActorRestored {
				conn.Close()
				logger.Info("I'm restored")
			}

			return nil
		}

		incomePacket := message.(IncomePacket)
		ctx := context.Background()

		if _, ok := packets.OpcodesToName[packets.Opcode(incomePacket.Opcode)]; ok {
			logger.Info("Income packet", "opcode", packets.OpcodesToName[packets.Opcode(incomePacket.Opcode)])
		}

		var respPkt packets.PacketEncodeInterface

		switch packets.Opcode(incomePacket.Opcode) {
		case packets.OpcodeAuth:
			var err error
			respPkt, err = handleAuth(ctx, &tempData, incomePacket.Data)
			if err != nil {
				return err
			}

			logger.Info("Active login", "activeLogin", tempData.ActiveLogin)
		case packets.OpcodeCreateCharacter:
			var err error
			respPkt, err = handleCreateCharacter(ctx, logger, &tempData, incomePacket.Data)
			if err != nil {
				return err
			}
		case packets.OpcodeRemoveCharacter:
			var err error
			respPkt, err = handleRemoveCharacter(ctx, logger, &tempData, incomePacket.Data)
			if err != nil {
				return err
			}
		case packets.OpcodeCreatePincode:
			var err error
			respPkt, err = handleCreatePincode(ctx, logger, &tempData, incomePacket.Data)
			if err != nil {
				return err
			}
		case packets.OpcodeChangePincode:
			var err error
			respPkt, err = handleUpdatePincode(ctx, logger, &tempData, incomePacket.Data)
			if err != nil {
				return err
			}
		case packets.OpcodeEnterGame:
			var err error
			respPkt, err = handleEnterGame(ctx, logger, &tempData, incomePacket.Data)
			if err != nil {
				return err
			}
		case packets.OpcodeExit:
			conn.Close()
			return nil
		default:
			packets.PrintFormattedHex(incomePacket.Data)
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

func handleAuth(ctx context.Context, tempData *TempData, data []byte) (packets.PacketEncodeInterface, error) {
	authPkt := &packets.Auth{}

	err := authPkt.Decode(ctx, bytes.NewReader(data), binary.BigEndian)
	if err != nil {
		return nil, err
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

	tempData.ActiveLogin = authPkt.Login

	cs := packets.NewCharacterScreen()

	characters := actor.SendToMultipleAndGetOne(actor.ActorRegistry.GetByName("storage"), GetCharacters{
		Login: authPkt.Login,
	}).([]packets.Character)

	cs.CharacterLen = uint8(len(characters))
	cs.Characters = characters

	if len(account.PincodeHash) != 0 {
		cs.Pincode = 1
	}

	return cs, nil
}

func handleCreateCharacter(ctx context.Context, logger *slog.Logger, tempData *TempData, data []byte) (packets.PacketEncodeInterface, error) {
	createCharPkt := packets.NewCharacterCreate()

	err := createCharPkt.Decode(ctx, bytes.NewReader(data), binary.BigEndian)
	if err != nil {
		return nil, err
	}

	logger.Info("Create character", "name", createCharPkt.Name)

	actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), SaveCharacter{
		Login: tempData.ActiveLogin,
		Data: packets.Character{
			IsActive: true,
			Name:     createCharPkt.Name,
			Job:      "Newbie",
			Level:    1,
			LookSize: createCharPkt.LookSize,
			Look:     createCharPkt.Look,
		},
	})

	return packets.NewCharacterCreateReply(), nil
}

func handleRemoveCharacter(ctx context.Context, logger *slog.Logger, tempData *TempData, data []byte) (packets.PacketEncodeInterface, error) {
	pkt := packets.NewCharacterRemove()

	err := pkt.Decode(ctx, bytes.NewReader(data), binary.BigEndian)
	if err != nil {
		return nil, err
	}

	logger.Info("Remove character", "pkt", pkt)

	actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), RemoveCharacter{
		Login: tempData.ActiveLogin,
		Name:  pkt.Name,
	})

	return packets.NewCharacterRemoveReply(), nil
}

func handleCreatePincode(ctx context.Context, logger *slog.Logger, tempData *TempData, data []byte) (packets.PacketEncodeInterface, error) {
	pkt := packets.NewCreatePincode()

	err := pkt.Decode(ctx, bytes.NewReader(data), binary.BigEndian)
	if err != nil {
		return nil, err
	}

	logger.Info("Create pincode", "pkt", pkt)

	actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), UpdatePincode{
		Login: tempData.ActiveLogin,
		Hash:  pkt.Hash,
	})

	return packets.NewCreatePincodeReply(), nil
}

func handleUpdatePincode(ctx context.Context, logger *slog.Logger, tempData *TempData, data []byte) (packets.PacketEncodeInterface, error) {
	pkt := packets.NewUpdatePincode()

	err := pkt.Decode(ctx, bytes.NewReader(data), binary.BigEndian)
	if err != nil {
		return nil, err
	}

	logger.Info("Update pincode", "pkt", pkt)

	actor.SendToMultiple(actor.ActorRegistry.GetByName("storage"), UpdatePincode{
		Login: tempData.ActiveLogin,
		Hash:  pkt.Hash,
	})

	return packets.NewUpdatePincodeReply(), nil
}

func handleEnterGame(ctx context.Context, logger *slog.Logger, tempData *TempData, data []byte) (packets.PacketEncodeInterface, error) {
	worldID := uint32(27012)

	lookItems := [10]packets.CharacterLookItem{}

	kitbagItems := []packets.KitbagItem{}
	for i := range 24 {
		kitbagItems = append(kitbagItems, packets.KitbagItem{
			GridID: uint16(i),
			ID:     0,
		})
	}
	kitbagItems = append(kitbagItems, packets.KitbagItem{
		GridID: uint16(65_535),
		ID:     0,
	})

	enterGameReply := packets.EnterGame{
		EnterRet:   0,
		AutoLock:   0,
		KitbagLock: 0,
		EnterType:  0,
		IsNewChar:  0,
		MapName:    "garner",
		CanTeam:    1,

		CharacterBase: packets.CharacterBase{
			ChaID:      4,
			WorldID:    worldID,
			CommID:     worldID,
			CommName:   "",
			GmLvl:      0,
			Handle:     33_565_845,
			CtrlType:   1,
			Name:       "name",
			MottoName:  "",
			Icon:       4,
			GuildID:    0,
			GuildName:  "",
			GuildMotto: "",
			StallName:  "",
			State:      0,
			Position: packets.Position{
				X:      217_475,
				Y:      278_175,
				Radius: 40,
			},
			Angle:        180,
			TeamLeaderID: 0,
			Side: packets.CharacterSide{
				SideID: 0,
			},
			Look: packets.CharacterLook{
				SynType: 0,
				TypeID:  2,
				IsBoat:  0,
				LookHuman: packets.CharacterLookHuman{
					HairID:   0,
					ItemGrid: lookItems,
				},
			},
			PkCtrl: packets.CharacterPK{
				PkCtrl: 0,
			},
			LookAppend: [4]packets.CharacterAppendLook{},
		},

		CharacterSkillBag: packets.CharacterSkillBag{
			SkillID:  36,
			Type:     0,
			SkillNum: 0,
			Skills:   []packets.CharacterSkill{},
		},

		CharacterSkillState: packets.CharacterSkillState{
			StatesLen: 0,
			States:    []packets.SkillState{},
		},

		CharacterAttribute: packets.CharacterAttribute{
			Type:       0,
			Num:        0,
			Attributes: []packets.Attribute{},
		},

		CharacterKitbag: packets.CharacterKitbag{
			Type:      0,
			KeybagNum: uint16(len(kitbagItems)),
			Items:     kitbagItems,
		},

		CharacterShortcut: packets.CharacterShortcut{
			Shortcuts: [36]packets.Shortcut{},
		},

		CharacterBoats: []packets.CharacterBoat{},

		ChaMainID: worldID,
	}

	return &enterGameReply, nil
}
