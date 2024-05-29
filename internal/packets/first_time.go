package packets

import (
	"fmt"
	"time"
)

//go:generate diho_bytes_generate first_time.go packets
type FirstTime struct {
	Header `dbg:"ignore"`
	Time   string
}

func (f FirstTime) Opcode() uint16 {
	return 940
}

func NewFirstTime() *FirstTime {
	timeNow := time.Now()

	return &FirstTime{
		Time: fmt.Sprintf("[%02d-%02d %02d:%02d:%02d:%03d]", timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond()/1000000),
	}
}
