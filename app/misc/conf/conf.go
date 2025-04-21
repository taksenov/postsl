/*
Copyright © 2025 taksenov@gmail.com
*/

// Package conf -- Application config.
package conf

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ServerAddr struct {
	Name string `mapstructure:"name"`
	Addr string `mapstructure:"addr"`
}

type Command struct {
	Command    string `mapstructure:"command"`
	IsTemplate bool   `mapstructure:"is_template"`
}

type AppConfig struct {
	Concurency struct {
		Use   bool `mapstructure:"use"`
		Coeff int  `mapstructure:"coeff"`
	} `mapstructure:"concurency"`

	Servers struct {
		Use  string       `mapstructure:"use"`
		List []ServerAddr `mapstructure:"list"`
	} `mapstructure:"servers"`

	RemoteCommand []Command `mapstructure:"remote_command"`

	Command struct {
		Main   string   `mapstructure:"main"`
		Params []string `mapstructure:"params"`
	} `mapstructure:"command"`
}

func LoadConfig(configFile string) (*AppConfig, error) {
	_, err := os.Stat(configFile)
	if err == nil {
		viper.SetConfigFile(configFile)
		viper.SetConfigType("yaml")
	} else {
		log.Fatal("cannot load config:", err)
	}

	config := &AppConfig{}

	viper.SetEnvPrefix("postsl")
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into config struct, %v", err)
	}

	return config, nil
}
