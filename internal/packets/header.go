package packets

import (
	"encoding/binary"

	utils "github.com/Nyarum/diho_bytes_generate/utils"
	bytebufferpool "github.com/valyala/bytebufferpool"
)

type PacketEncodeInterface interface {
	SetHeaderLen(len uint16)
	EncodeHeader(endian binary.ByteOrder) ([]byte, error)
	Encode(endian binary.ByteOrder) ([]byte, error)
}

type Header struct {
	Len uint16
	ID  uint32
}

func (h *Header) SetHeaderLen(len uint16) {
	h.Len = len
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
	return utils.Clone(newBuf), nil
}

func EncodeWithHeader(pkt PacketEncodeInterface, endian binary.ByteOrder) ([]byte, error) {
	bodyBuf, err := pkt.Encode(endian)
	if err != nil {
		return nil, err
	}

	pkt.SetHeaderLen(uint16(len(bodyBuf)) + 6)

	headerBuf, err := pkt.EncodeHeader(endian)
	if err != nil {
		return nil, err
	}

	return append(headerBuf, bodyBuf...), nil
}
