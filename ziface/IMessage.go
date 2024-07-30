package ziface

type IMessage interface {
	GetMessageId() uint32
	GetMessageLen() uint32
	GetData() []byte
	SetMessageId(id uint32)
	SetMessageLen(len uint32)
	SetData(data []byte)
}
