package config

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"sync"
)

var (
	loaderInstance *Configuration
	once           sync.Once
)

func LoadConfig() *Configuration {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.Set("Verbose", true)
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			panic("Error reading configuration: " + err.Error())
		}

		loaderInstance = &Configuration{}

		if err := viper.Unmarshal(loaderInstance); err != nil {
			panic("Error interpretando configuración: " + err.Error())
		}

		validate := validator.New()
		if err := validate.Struct(loaderInstance); err != nil {
			panic("Error validando configuración: " + err.Error())
		}
	})

	return loaderInstance
}
func GetConfig[T any](key string) (T, error) {
	config := LoadConfig()

	keys := strings.Split(key, ".")

	value := reflect.ValueOf(config).Elem()
	var finalValue reflect.Value
	for _, k := range keys {
		fieldVal := value.FieldByName(strings.Title(k))

		if !fieldVal.IsValid() {
			return zeroOf[T](), fmt.Errorf("config key %s not found", key)
		}

		if fieldVal.Kind() == reflect.Struct && k != keys[len(keys)-1] {
			value = fieldVal
		} else if k == keys[len(keys)-1] {
			finalValue = fieldVal
		} else {
			return zeroOf[T](), fmt.Errorf("config key %s not found", key)
		}
	}

	result, ok := finalValue.Interface().(T)
	if !ok {
		return zeroOf[T](), errors.New(fmt.Sprintf("type mismatch: expected %T, but found %T", result, finalValue.Interface()))
	}
	return result, nil
}

func zeroOf[T any]() T {
	var v T
	return v
}
