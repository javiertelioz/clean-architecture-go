package config

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
			panic("Error interpreting configuration: " + err.Error())
		}

		validate := validator.New()
		if err := validate.Struct(loaderInstance); err != nil {
			panic("Error validating configuration: " + err.Error())
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
		causer := cases.Title(language.English, cases.NoLower)
		fieldVal := value.FieldByName(causer.String(k))

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
		return zeroOf[T](), fmt.Errorf("type mismatch: expected %T, but found %T", result, finalValue.Interface())
	}

	return result, nil
}

func zeroOf[T any]() T {
	var v T
	return v
}
