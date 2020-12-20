package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

// ErrHeader Head Error
var ErrHeader = errors.New("Header error")

// MsgHeader 消息头
const MsgHeader = "-01-"

// Encode 对数据进行编码 |-- Header --|-- contentlen --|-- content --|
func Encode(bytesBuffer io.Writer, content string) error {
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(MsgHeader)); err != nil {
		return err
	}
	clen := uint32(len(content))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

// Decode 对数据进行解码，获取具体数据
func Decode(bytesBuffer io.Reader) (bodyBuffer []byte, err error) {
	magicBuffer := make([]byte, len(MsgHeader))
	if _, err := io.ReadFull(bytesBuffer, magicBuffer); err != nil {
		return nil, err
	}
	if string(magicBuffer) != MsgHeader {
		return nil, ErrHeader
	}

	lengthBuffer := make([]byte, 4) // 4  sizeof(int32)
	if _, err := io.ReadFull(bytesBuffer, lengthBuffer); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuffer)
	bodyBuffer = make([]byte, length)
	if _, err := io.ReadFull(bytesBuffer, bodyBuffer); err != nil {
		return nil, err
	}
	return bodyBuffer, err
}
