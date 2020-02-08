package example

import (
	"fmt"
	"testing"
)
import "gopkg.in/ini.v1"

func TestReadConfig(t *testing.T) {

	cfg, err := ini.Load("../config.ini")
	if err != nil {
		fmt.Println("文件加载失败...S")
		return
	}

	port := cfg.Section("app").Key("server.port").MustInt(9999)
	fmt.Println("加载配置文件端口:", port)
	fmt.Println("应用名称: ", cfg.Section("app").Key("server.name").MustString(""))

}
