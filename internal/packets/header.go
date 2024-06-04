package packets

import (
	"bytes"
	"context"
	"encoding/binary"

	utils "github.com/Nyarum/diho_bytes_generate/utils"
	bytebufferpool "github.com/valyala/bytebufferpool"
)

type PacketEncodeInterface interface {
	Opcode() uint16
	SetHeader(len, opcode uint16)
	EncodeHeader(endian binary.ByteOrder) ([]byte, error)
	Encode(ctx context.Context, endian binary.ByteOrder) ([]byte, error)
}

type Header struct {
	Len    uint16
	ID     uint32
	Opcode uint16
}

func (h *Header) SetHeader(len, opcode uint16) {
	h.Len = len
	h.Opcode = opcode
}

func (h Header) EncodeHeader(endian binary.ByteOrder) ([]byte, error) {
	newBuf := bytebufferpool.Get()
	defer bytebufferpool.Put(newBuf)
	var err error
	err = binary.Write(newBuf, endian, h.Len)
	if err != nil {
		return nil, err
	}

	h.ID = 128
	err = binary.Write(newBuf, binary.LittleEndian, h.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(newBuf, endian, h.Opcode)
	if err != nil {
		return nil, err
	}

	return utils.Clone(newBuf), nil
}

func EncodeWithHeader(ctx context.Context, pkt PacketEncodeInterface, endian binary.ByteOrder) ([]byte, error) {
	bodyBuf, err := pkt.Encode(ctx, endian)
	if err != nil {
		return nil, err
	}

	pkt.SetHeader(uint16(len(bodyBuf))+8, pkt.Opcode())

	headerBuf, err := pkt.EncodeHeader(endian)
	if err != nil {
		return nil, err
	}

	return append(headerBuf, bodyBuf...), nil
}

func DecodeHeader(buf []byte) (*Header, error) {
	reader := bytes.NewReader(buf)
	var header Header

	err := binary.Read(reader, binary.BigEndian, &header.Len)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.LittleEndian, &header.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.BigEndian, &header.Opcode)
	if err != nil {
		return nil, err
	}

	return &header, nil
}
