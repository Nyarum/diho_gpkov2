package packets

import "context"

const (
	SynLookSwitch uint8 = iota
	SynLookChange
)

const (
	SYN_KITBAG_INIT uint8 = iota
	SYN_KITBAG_EQUIP
	SYN_KITBAG_UNFIX
	SYN_KITBAG_PICK
	SYN_KITBAG_THROW
	SYN_KITBAG_SWITCH
	SYN_KITBAG_TRADE
	SYN_KITBAG_FROM_NPC
	SYN_KITBAG_TO_NPC
	SYN_KITBAG_SYSTEM
	SYN_KITBAG_FORGES
	SYN_KITBAG_FORGEF
	SYN_KITBAG_BANK
	SYN_KITBAG_ATTR
)

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
	GridID        uint16
	ID            uint16
	Num           uint16
	Endure        [2]uint16
	Energy        [2]uint16
	ForgeLevel    uint8
	IsValid       bool
	ItemDBInstID  uint32 `dbg:"ID==3988"`
	ItemDBForge   uint32
	BoatNull      uint32 `dbg:"ID==3988"`
	ItemDBInstID2 uint32 `dbg:"ID!=3988"`
	IsParams      bool
	InstAttrs     [5]InstAttr
}

func (p *KitbagItem) Filter(ctx context.Context, name string) bool {
	switch name {
	case "GridID":
		return p.GridID == 65535
	case "ID":
		return p.ID == 0
	case "IsParams":
		return !p.IsParams
	}

	return false
}

type CharacterKitbag struct {
	Type      uint8
	KeybagNum uint16 `dbg:"Type==SYN_KITBAG_INIT"`
	Items     []KitbagItem
}

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
	SynType     uint8 `dbg:"ignore"`
	ID          uint16
	ItemSync    CharacterLookItemSync `dbg:"SynType==SynLookChange"`
	ItemShow    CharacterLookItemShow `dbg:"SynType==SynLookSwitch"`
	IsDBParams  uint8
	DBParams    [2]uint32
	IsInstAttrs uint8
	InstAttrs   [5]InstAttr
}

func (p *CharacterLookItem) Filter(ctx context.Context, name string) bool {
	switch name {
	case "ID":
		return p.ID == 0
	case "IsDBParams":
		return p.IsDBParams == 0
	case "IsInstAttrs":
		return p.IsInstAttrs == 0
	}

	return false
}

type CharacterLookHuman struct {
	HairID   uint16
	ItemGrid [10]CharacterLookItem
}

type CharacterLook struct {
	SynType   uint8
	TypeID    uint16
	IsBoat    uint8
	LookBoat  CharacterLookBoat  `dbg:"IsBoat==1"`
	LookHuman CharacterLookHuman `dbg:"IsBoat==0"`
}

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
