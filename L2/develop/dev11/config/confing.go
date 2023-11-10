package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type MyConfig struct {
	Host string ` yaml:"host" env-default:"localhost"`
	Port string ` yaml:"port" env-default:"8080"`
}

var instance *MyConfig
var once sync.Once

// патерн синглтон
// поулчаем данные с конфиг файла
func GetConfig() *MyConfig {
	once.Do(func() {
		instance = &MyConfig{}
		if err := cleanenv.ReadConfig("../config/config.yml", instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
