// Code generated by diho_bytes_generate auth.go; DO NOT EDIT.

package packets

import (
	"context"
	"encoding/binary"
	utils "github.com/Nyarum/diho_bytes_generate/utils"
	bytebufferpool "github.com/valyala/bytebufferpool"
)

func (p *Auth) Encode(ctx context.Context, endian binary.ByteOrder) ([]byte, error) {
	newBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(newBuf)
	var err error
	err = utils.WriteBytes(newBuf, p.Key)
	if err != nil {
		return nil, err
	}
	err = utils.WriteStringNull(newBuf, p.Login)
	if err != nil {
		return nil, err
	}
	err = utils.WriteBytes(newBuf, p.Password)
	if err != nil {
		return nil, err
	}
	err = utils.WriteStringNull(newBuf, p.MAC)
	if err != nil {
		return nil, err
	}
	err = binary.Write(newBuf, endian, p.IsCheat)
	if err != nil {
		return nil, err
	}
	err = binary.Write(newBuf, endian, p.ClientVersion)
	if err != nil {
		return nil, err
	}
	return utils.Clone(newBuf), nil
}
