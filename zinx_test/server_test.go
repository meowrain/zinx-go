package zinxtest

import (
	"testing"
	"zinx/znet"
)

func TestServer(t *testing.T) {
	t.Log("TestServer")
	s := znet.NewServer("meowrain server")
	s.Serve()
}
