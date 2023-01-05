package app

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

//https://juejin.cn/post/7016979344094396424
const (
	defaultConfigFileName = "config/application.yaml"
	EnvKeyConfig          = "config"
)

type server struct {
	Port int32 `mapstructure:"port"`
}

type datasource struct {
	Dsn         string `mapstructure:"dsn"`
	AutoMigrate bool   `mapstructure:"auto-migrate"`
}

type config struct {
	Server     server     `mapstructure:"server"`
	Datasource datasource `mapstructure:"datasource"`
}

func initConfig(app *application) {
	var configFile string
	if configEnv := os.Getenv(EnvKeyConfig); configEnv != "" {
		configFile = configEnv
	} else {
		workDir, _ := os.Getwd()
		configFile = path.Join(workDir, defaultConfigFileName)
	}
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config error %v", err))
	}
	if err := v.Unmarshal(app.Config); err != nil {
		panic(fmt.Errorf("config mapping error %v", err))
	}
}
