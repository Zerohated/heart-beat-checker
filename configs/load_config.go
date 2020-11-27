package configs

import (
	"log"

	"github.com/BurntSushi/toml"
)

func init() {
	_, err := toml.DecodeFile("./configs/config.toml", &Config)
	// _, err := toml.DecodeFile("/mnt/c/Users/David/workspace/xmas2019/config/test_config.toml", &Config)
	if err != nil {
		log.Println(err.Error())
	}
}
