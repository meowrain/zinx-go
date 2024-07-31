package zinxtest

import (
	"net"
	"testing"
	"zinx/znet"
)

func TestDataPack(t *testing.T) {
	t.Log("TestDataPack")
	datapack := znet.NewDataPack()
	msg := new(znet.Message)
	str := "helloworld,meowrain"

	msg.SetMessageLen(uint32(len(str)))
	msg.SetMessageId(1)

	msg.SetData([]byte(str))
	pack, err := datapack.Pack(msg)
	if err != nil {
		t.Error(err)
	}
	t.Logf("pack:%v\n", pack)
	msg2, err := datapack.Unpack(pack)
	if err != nil {
		t.Error(err)
	}
	t.Logf("unpack msg:%v\n", msg2)

}
func TestDataPackClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer conn.Close()
	datapack := znet.NewDataPack()
	msg := new(znet.Message)
	str := "helloworld,meowrain"
	msg.SetMessageLen(uint32(len(str)))
	msg.SetMessageId(1)
	msg.SetData([]byte(str))
	pack, err := datapack.Pack(msg)
	if err != nil {
		t.Error(err)
	}
	t.Logf("pack:%v\n", pack)
	_, err = conn.Write(pack)
	if err != nil {
		t.Error(err)
	}

}
