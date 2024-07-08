package zutils

import (
	"fmt"
	"os"
	"zinx/ziface"

	"gopkg.in/yaml.v3"
)

type GlobalObj struct {
	/*Server*/
	TcpServer ziface.IServer
	Host      string
	Port      int
	Name      string

	/*Zinx*/
	Version         string
	MaxConnection   int
	MaxPackagesSize uint32
}

var GlobalObject *GlobalObj = &GlobalObj{}

func (g *GlobalObj) String() string {
	return fmt.Sprintf(
		"GlobalObj{Name: %s, Host: %s, Port: %d, Version: %s, MaxConnection: %d, MaxPackagesSize: %d}",
		g.Name, g.Host, g.Port, g.Version, g.MaxConnection, g.MaxPackagesSize,
	)
}
func (g *GlobalObj) LoadFromConfigFile(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &Cfg); err != nil {
		return err
	}
	return nil
}
