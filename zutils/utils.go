package zutils

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	App struct {
		Name            string `yaml:"name"`
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		Version         string `yaml:"version"`
		MaxConnection   int    `yaml:"max_connection"`
		MaxPackagesSize uint32 `yaml:"max_packages_size"`
	} `yaml:"app"`
}

var Cfg Config = Config{}

func init() {
	if err := GlobalObject.LoadFromConfigFile("config/config.yaml"); err != nil {
		logrus.Errorln(err)
		//如果配置文件加载失败，设置默认值
		GlobalObject = &GlobalObj{
			TcpServer:       nil,
			Name:            "ZinxServer",
			Host:            "0.0.0.0",
			Port:            8080,
			Version:         "tcp4",
			MaxConnection:   1000,
			MaxPackagesSize: 1000,
		}
		return
	}

	GlobalObject = &GlobalObj{
		Name:            Cfg.App.Name,
		Host:            Cfg.App.Host,
		Port:            Cfg.App.Port,
		Version:         Cfg.App.Version,
		MaxConnection:   Cfg.App.MaxConnection,
		MaxPackagesSize: Cfg.App.MaxPackagesSize,
	}

}
