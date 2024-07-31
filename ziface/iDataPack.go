package ziface

type IDataPack interface {
	// 获取头部商都
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//解包方法
	Unpack([]byte) (IMessage, error)
	//解包头方法
	UnpackHead([]byte) (IMessage, error)
}
