package app

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
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

type auth struct {
	Excludes []string `mapstructure:"excludes"`
	Token    token    `mapstructure:"token"`
}

type token struct {
	ExpireTime time.Duration `mapstructure:"expire-time"`
	SigningKey string        `mapstructure:"signing-key"`
}

type datavcs struct {
	RootPath string `mapstructure:"root-path"`
}

type config struct {
	Server     server     `mapstructure:"server"`
	Datasource datasource `mapstructure:"datasource"`
	Auth       auth       `mapstructure:"auth"`
	Datavcs    datavcs    `mapstructure:"datavcs"`
}

func initConfig(app *Application) {
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
