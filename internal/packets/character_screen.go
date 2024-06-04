package packets

//go:generate diho_bytes_generate character_screen.go
type CharacterScreen struct {
	Header       `dbg:"ignore"`
	ErrorCode    uint16
	Key          []byte
	CharacterLen uint8
	//Characters   []Character
	Pincode    uint8
	Encryption uint32
	DWFlag     uint32
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
