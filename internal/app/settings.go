package app

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

const (
	tagName        = "env"
	settingsPrefix = "URL_SHORT"
)

// settings holds the application main settings
type settings struct {
	port          int    `env:"PORT"`
	dbURL         string `env:"DATABASE_URL"`
	redisAddr     string `env:"REDIS_ADDR"`
	redisPassword string `env:"REDIS_PASSWORD"`
}

// loadSettings loads application main settings
// Output:
//    settings in case of success, otherwise returns error
func loadSettings() (*settings, error) {
	viper.SetEnvPrefix(settingsPrefix)
	s := &settings{}

	// bind values to viper
	if err := s.forAll(func(tag string, isZero bool) error {
		return viper.BindEnv(tag)
	}); err != nil {
		return nil, err
	}

	s.port = viper.GetInt("port")
	s.dbURL = viper.GetString("database_url")
	s.redisAddr = viper.GetString("redis_addr")
	s.redisPassword = viper.GetString("redis_password")

	// check if all parametres are set
	err := s.forAll(func(tag string, isZero bool) error {
		if isZero {
			return fmt.Errorf("environment value `%s_%s` should be set", settingsPrefix, tag)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return s, nil
}

// forAll applies provided function to all fields of settings struct
// Inputs:
//   f function than needs to be applied
// Output:
//   returns error if provided functions returns error
func (s *settings) forAll(f func(tag string, isZero bool) error) error {
	t := reflect.TypeOf(*s)
	v := reflect.ValueOf(*s)
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if err := f(field.Tag.Get(tagName), v.Field(i).IsZero()); err != nil {
			return err
		}
	}

	return nil
}
