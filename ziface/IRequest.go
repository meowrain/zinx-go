package ziface

type IRequest interface {
	//得到当前链接
	GetConnection() IConnection

	//得到当前数据
	GetData() []byte
}
