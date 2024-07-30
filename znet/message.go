package znet

type Message struct {
	id      uint32 //消息id
	datalen uint32 //消息长度
	data    []byte //消息内容
}

func (m *Message) GetMessageId() uint32 {
	return m.id
}
func (m *Message) GetMessageLen() uint32 {
	return m.datalen
}
func (m *Message) GetData() []byte {
	return m.data
}
func (m *Message) SetMessageId(id uint32) {
	m.id = id
}
func (m *Message) SetMessageLen(len uint32) {
	m.datalen = len
}
func (m *Message) SetData(data []byte) {
	m.data = data
}
