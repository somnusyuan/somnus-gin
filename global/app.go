package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"somnus-gin/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
}

var App = new(Application)
