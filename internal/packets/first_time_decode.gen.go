// Code generated by diho_bytes_generate first_time.go; DO NOT EDIT.

package packets

import (
	"context"
	"encoding/binary"
	utils "github.com/Nyarum/diho_bytes_generate/utils"
	"io"
)

func (p *FirstTime) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	p.Time, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	return nil
}
