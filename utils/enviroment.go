package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var Env *viper.Viper

func init() {
	Env = viper.New()
}

func ParseEnvFile(filename, filetype string, paths ...string) error {

	Env.SetConfigName(filename)
	Env.SetConfigType(filetype)

	for _, path := range paths {
		Env.AddConfigPath(path)
	}

	if err := Env.ReadInConfig(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed parsing env file %s.%s", filename, filetype))
	}

	return nil
}