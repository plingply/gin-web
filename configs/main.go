package configs

import (
	"log"

	"github.com/go-ini/ini"
)

var Cfg, _ = ini.Load("configs/app.ini")

func Get(key string) *ini.Section {
	sec, err := Cfg.GetSection(key)
	if err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}
	return sec
}
