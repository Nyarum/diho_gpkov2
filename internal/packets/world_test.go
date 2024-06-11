package packets

import (
	"context"
	"encoding/binary"
	"testing"
)

func TestWorldEncode(t *testing.T) {
	worldID := uint32(27012)

	lookItems := [10]CharacterLookItem{}

	kitbagItems := []KitbagItem{}
	for i := range 24 {
		kitbagItems = append(kitbagItems, KitbagItem{
			GridID: uint16(i),
			ID:     0,
		})
	}
	kitbagItems = append(kitbagItems, KitbagItem{
		GridID: uint16(65_535),
		ID:     0,
	})

	enterGameReply := EnterGame{
		EnterRet:   0,
		AutoLock:   0,
		KitbagLock: 0,
		EnterType:  0,
		IsNewChar:  0,
		MapName:    "garner",
		CanTeam:    1,

		CharacterBase: CharacterBase{
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
			Position: Position{
				X:      217_475,
				Y:      278_175,
				Radius: 40,
			},
			Angle:        180,
			TeamLeaderID: 0,
			Side: CharacterSide{
				SideID: 0,
			},
			Look: CharacterLook{
				SynType: 0,
				TypeID:  2,
				IsBoat:  0,
				LookHuman: CharacterLookHuman{
					HairID:   0,
					ItemGrid: lookItems,
				},
			},
			PkCtrl: CharacterPK{
				PkCtrl: 0,
			},
			LookAppend: [4]CharacterAppendLook{},
		},

		CharacterSkillBag: CharacterSkillBag{
			SkillID:  36,
			Type:     0,
			SkillNum: 0,
			Skills:   []CharacterSkill{},
		},

		CharacterSkillState: CharacterSkillState{
			StatesLen: 0,
			States:    []SkillState{},
		},

		CharacterAttribute: CharacterAttribute{
			Type:       0,
			Num:        0,
			Attributes: []Attribute{},
		},

		CharacterKitbag: CharacterKitbag{
			Type:      0,
			KeybagNum: uint16(len(kitbagItems)),
			Items:     kitbagItems,
		},

		CharacterShortcut: CharacterShortcut{
			Shortcuts: [36]Shortcut{},
		},

		CharacterBoats: []CharacterBoat{},

		ChaMainID: worldID,
	}

	buf, _ := enterGameReply.Encode(context.Background(), binary.BigEndian)
	_ = buf
}
