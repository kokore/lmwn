package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type ConfigType struct {
	App      AppConfig
	CovidAPI CovidAPIConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type CovidAPIConfig struct {
	CovidCaseURL string
}

var Config *ConfigType

func InitConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.ReadInConfig()

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
