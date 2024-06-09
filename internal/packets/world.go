package packets

//go:generate diho_bytes_generate world.go
type CharacterBoat struct {
	CharacterBase       CharacterBase
	CharacterAttribute  CharacterAttribute
	CharacterKitbag     CharacterKitbag
	CharacterSkillState CharacterSkillState
}

type Shortcut struct {
	Type   uint8
	GridID uint16
}

type CharacterShortcut struct {
	Shortcuts [36]Shortcut
}

type KitbagItem struct {
	GridID       uint16
	ID           uint16
	Num          uint16
	Endure       [2]uint16
	Energy       [2]uint16
	ForgeLevel   uint8
	IsValid      bool
	ItemDBInstID uint32
	ItemDBForge  uint32
	IsParams     bool
	InstAttrs    [5]InstAttr
}

/*
func (k *KitbagItem) Process(buf *[]byte, mode ...processor.Mode) {
	defaultMode := processor.Write
	if len(mode) > 0 {
		defaultMode = mode[0]
	}

	p := processor.NewProcessor(defaultMode)
	p.UInt16(buf, &k.GridID)

	if k.GridID.Value == 65535 {
		return
	}

	p.UInt16(buf, &k.ID)

	if k.ID.Value > 0 {
		p.UInt16(buf, &k.Num)

		for v := range k.Endure {
			p.UInt16(buf, &k.Endure[v])
		}

		for v := range k.Energy {
			p.UInt16(buf, &k.Energy[v])
		}

		p.UInt8(buf, &k.ForgeLevel)
		p.Bool(buf, &k.IsValid)

		//if "item_info.type" == "boat" {
		if k.ID.Value == 3988 {
			p.UInt32(buf, &k.ItemDBInstID)
		}

		p.UInt32(buf, &k.ItemDBForge)

		//if "item_info.type" == "boat" {
		if k.ID.Value == 3988 {
			v := uint32{}
			p.UInt32(buf, &v)
		} else {
			p.UInt32(buf, &k.ItemDBInstID)
		}

		p.Bool(buf, &k.IsParams)

		if k.IsParams.Value {
			for ki := range k.InstAttrs {
				k.InstAttrs[ki].Process(buf, mode...)
			}
		}
	}
}
*/

type CharacterKitbag struct {
	Type      uint8
	KeybagNum uint16
	Items     []KitbagItem
}

/*
func (c *CharacterKitbag) Process(buf *[]byte, mode ...processor.Mode) {
	defaultMode := processor.Write
	if len(mode) > 0 {
		defaultMode = mode[0]
	}

	p := processor.NewProcessor(defaultMode)
	p.UInt8(buf, &c.Type)

	if c.Type.Value == types.SYN_KITBAG_INIT {
		p.UInt16(buf, &c.KeybagNum)
	}

	c.KeybagNum.Value = c.KeybagNum.Value + 1

	if len(c.Items) != int(c.KeybagNum.Value) {
		c.Items = make([]KitbagItem, c.KeybagNum.Value)
	}

	for k := range c.Items {
		c.Items[k].Process(buf, mode...)
	}
}
*/

type Attribute struct {
	ID    uint8
	Value uint32
}

type CharacterAttribute struct {
	Type       uint8
	Num        uint16
	Attributes []Attribute
}

type SkillState struct {
	ID    uint8
	Level uint8
}

type CharacterSkillState struct {
	StatesLen uint8
	States    []SkillState
}

type CharacterSkill struct {
	ID         uint16
	State      uint8
	Level      uint8
	UseSP      uint16
	UseEndure  uint16
	UseEnergy  uint16
	ResumeTime uint32
	RangeType  uint16
	Params     []uint16 // ?
}

type CharacterSkillBag struct {
	SkillID  uint16
	Type     uint8
	SkillNum uint16
	Skills   []CharacterSkill
}

type CharacterAppendLook struct {
	LookID  uint16
	IsValid uint8
}

type CharacterPK struct {
	PkCtrl uint8
}

type CharacterLookBoat struct {
	PosID     uint16
	BoatID    uint16
	Header    uint16
	Body      uint16
	Engine    uint16
	Cannon    uint16
	Equipment uint16
}

type CharacterLookItemSync struct {
	Endure  uint16
	Energy  uint16
	IsValid uint8
}

type CharacterLookItemShow struct {
	Num        uint16
	Endure     [2]uint16
	Energy     [2]uint16
	ForgeLevel uint8
	IsValid    uint8
}

type CharacterLookItem struct {
	ID          uint16
	ItemSync    CharacterLookItemSync
	ItemShow    CharacterLookItemShow
	IsDBParams  uint8
	DBParams    [2]uint32
	IsInstAttrs uint8
	InstAttrs   [5]InstAttr
}

/*
func (c *CharacterLookItem) Process(buf *[]byte, synType uint8, mode ...processor.Mode) {
	defaultMode := processor.Write
	if len(mode) > 0 {
		defaultMode = mode[0]
	}

	p := processor.NewProcessor(defaultMode)
	p.UInt16(buf, &c.ID)

	if c.ID.Value != 0 {
		if synType.Value == types.SynLookChange {
			c.ItemSync.Process(buf, mode...)
		} else {
			c.ItemShow.Process(buf, mode...)
			p.UInt8(buf, &c.IsDBParams)

			if c.IsDBParams.Value != 0 {
				for k := range c.DBParams {
					p.UInt32(buf, &c.DBParams[k])
				}

				p.UInt8(buf, &c.IsInstAttrs)

				if c.IsInstAttrs.Value != 0 {
					for k := range c.InstAttrs {
						c.InstAttrs[k].Process(buf, mode...)
					}
				}
			}
		}
	}
}
*/

type CharacterLookHuman struct {
	HairID   uint16
	ItemGrid [10]CharacterLookItem
}

type CharacterLook struct {
	SynType   uint8
	TypeID    uint16
	IsBoat    uint8
	LookBoat  CharacterLookBoat  `dbg:"IsBoat=1"`
	LookHuman CharacterLookHuman `dbg:"IsBoat=0"`
}

/*
func (c *CharacterLook) Process(buf *[]byte, mode ...processor.Mode) {
	defaultMode := processor.Write
	if len(mode) > 0 {
		defaultMode = mode[0]
	}

	p := processor.NewProcessor(defaultMode)
	p.UInt8(buf, &c.SynType)
	p.UInt16(buf, &c.TypeID)
	p.UInt8(buf, &c.IsBoat)

	if c.IsBoat.Value == 1 {
		(&c.LookBoat).Process(buf, c.SynType, mode...)
	} else {
		(&c.LookHuman).Process(buf, c.SynType, mode...)
	}
}
*/

type EntityEvent struct {
	EntityID   uint32
	EntityType uint8
	EventID    uint16
	EventName  string
}

type CharacterSide struct {
	SideID uint8
}

type Position struct {
	X      uint32
	Y      uint32
	Radius uint32
}

type CharacterBase struct {
	ChaID        uint32
	WorldID      uint32
	CommID       uint32
	CommName     string
	GmLvl        uint8
	Handle       uint32
	CtrlType     uint8
	Name         string
	MottoName    string
	Icon         uint16
	GuildID      uint32
	GuildName    string
	GuildMotto   string
	StallName    string
	State        uint16
	Position     Position
	Angle        uint16
	TeamLeaderID uint32
	Side         CharacterSide
	EntityEvent  EntityEvent
	Look         CharacterLook
	PkCtrl       CharacterPK
	LookAppend   [4]CharacterAppendLook
}

type EnterGame struct {
	Header              `dbg:"ignore"`
	EnterRet            uint16
	AutoLock            uint8
	KitbagLock          uint8
	EnterType           uint8
	IsNewChar           uint8
	MapName             string
	CanTeam             uint8
	CharacterBase       CharacterBase
	CharacterSkillBag   CharacterSkillBag
	CharacterSkillState CharacterSkillState
	CharacterAttribute  CharacterAttribute
	CharacterKitbag     CharacterKitbag
	CharacterShortcut   CharacterShortcut
	BoatLen             uint8
	CharacterBoats      []CharacterBoat
	ChaMainID           uint32
}

func (e EnterGame) Opcode() uint16 {
	return 516
}

func NewEnterGame() *EnterGame {
	return &EnterGame{}
}

type EnterGameRequest struct {
	CharacterName string
}
