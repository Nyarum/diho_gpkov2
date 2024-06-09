package packets

type Opcode uint16

const (
	OpcodeAuth            Opcode = 431
	OpcodeCreateCharacter Opcode = 435
	OpcodeRemoveCharacter Opcode = 436
	OpcodeCreatePincode   Opcode = 346
	OpcodeChangePincode   Opcode = 347
	OpcodeEnterGame       Opcode = 433
	OpcodeExit            Opcode = 432
)
