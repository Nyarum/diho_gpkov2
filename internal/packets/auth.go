package packets

//go:generate diho_bytes_generate auth.go
type Auth struct {
	Key           []byte
	Login         string
	Password      []byte
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
	PincodeHash   string `dbg:"ignore"`
}

func (a Auth) Opcode() uint16 {
	return 431
}
