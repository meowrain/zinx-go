package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/ziface"
	"zinx/zutils"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen + id len
	// datalen = 4
	// id len = 4
	// total = 8
	return 8
}

// datalen | msgID | data
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	// 1. data len
	if err := binary.Write(dataBuff, binary.BigEndian, uint32(msg.GetMessageLen())); err != nil {
		return nil, err
	}
	// 2. id
	if err := binary.Write(dataBuff, binary.BigEndian, uint32(msg.GetMessageId())); err != nil {
		return nil, err
	}
	// 3. data
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}
func (dp *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(data)

	msg := new(Message)
	var datalen uint32
	var id uint32
	var data_s []byte

	if err := binary.Read(reader, binary.BigEndian, &datalen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &id); err != nil {
		return nil, err
	}
	data_s = make([]byte, datalen) // Create a new slice with the same length as the data length.
	if err := binary.Read(reader, binary.BigEndian, &data_s); err != nil {
		return nil, err
	}
	msg.SetMessageLen(datalen)
	msg.SetMessageId(id)
	msg.SetData(data_s) // Set the data field of the Message object.
	if zutils.GlobalObject.MaxPackagesSize > 0 && datalen > zutils.GlobalObject.MaxPackagesSize {
		return nil, errors.New("ErrPackageTooLarge")
	}
	return msg, nil
}
func (dp *DataPack) UnpackHead(data []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(data)
	msg := new(Message)
	var datalen uint32
	var id uint32
	if err := binary.Read(reader, binary.BigEndian, &datalen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &id); err != nil {
		return nil, err
	}
	msg.SetMessageId(id)
	msg.SetMessageLen(datalen)
	return msg, nil
}
