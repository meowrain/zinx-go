package zinxtest

import (
	"log"
	"testing"
	"zinx/zutils"
)

func TestConfigRead(t *testing.T) {
	log.Println(zutils.Cfg)
}
