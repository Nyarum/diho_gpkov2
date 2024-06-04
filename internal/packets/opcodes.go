package packets

type Opcode uint16

const (
	OpcodeAuth            Opcode = 431
	OpcodeCreateCharacter Opcode = 435
	OpcodeExit            Opcode = 432
)
