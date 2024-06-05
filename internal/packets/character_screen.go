package packets

//go:generate diho_bytes_generate character_screen.go
type CharacterScreen struct {
	Header       `dbg:"ignore"`
	ErrorCode    uint16
	Key          []byte
	CharacterLen uint8
	Characters   [3]Character
	Pincode      uint8
	Encryption   uint32
	DWFlag       uint32
}

func NewCharacterScreen() *CharacterScreen {
	return &CharacterScreen{
		Key:     []byte{0x7C, 0x35, 0x09, 0x19, 0xB2, 0x50, 0xD3, 0x49},
		DWFlag:  12820,
		Pincode: 1,
	}
}

func (c CharacterScreen) Opcode() uint16 {
	return 931
}

type Character struct {
	IsActive bool
	Name     string
	Job      string
	Map      string
	Level    uint16
	LookSize uint16
	Look     Look `dbg:"little"`
}

type Look struct {
	Ver       uint16
	TypeID    uint16
	ItemGrids [10]ItemGrid
	Hair      uint16
}

type ItemGrid struct {
	ID        uint16
	Num       uint16
	Endure    [2]uint16
	Energy    [2]uint16
	ForgeLv   uint8
	DBParams  [2]uint32
	InstAttrs [5]InstAttr
	ItemAttrs [40]ItemAttr
	IsChange  bool
}

type ItemAttr struct {
	Attr   uint16
	IsInit bool
}

type InstAttr struct {
	ID    uint16
	Value uint16
}

type CharacterCreate struct {
	Name     string
	Map      string
	LookSize uint16
	Look     Look `dbg:"little"`
}

func NewCharacterCreate() *CharacterCreate {
	return &CharacterCreate{}
}

func (c CharacterCreate) Opcode() uint16 {
	return 435
}

type CharacterCreateReply struct {
	Header    `dbg:"ignore"`
	ErrorCode uint16
}

func (c CharacterCreateReply) Opcode() uint16 {
	return 935
}

func NewCharacterCreateReply() *CharacterCreateReply {
	return &CharacterCreateReply{}
}
