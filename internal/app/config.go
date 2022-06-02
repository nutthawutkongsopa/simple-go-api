package app

import (
	"fmt"
	"path/filepath"
	"test-api/internal/core"

	"github.com/spf13/viper"
)

type Environment struct {
	EnvName string `mapstructure:"TEST_API_ENVIRONMENT"`
}

type AppSettings struct {
	DBConnectionString string
	ApplicateionURL    string
	GINMode            string
	AppHost            string
}

func GetEnv() (*Environment, error) {
	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigFile("./configs/app.env")
	v.SetConfigType("env")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := Environment{}
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetAppSettigs() (*AppSettings, error) {
	env, err := GetEnv()
	if err != nil {
		return nil, err
	}
	rootDir, err := core.GetCurrentDirectory()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(rootDir, "configs")
	result := AppSettings{}
	viper := viper.New()
	viper.AddConfigPath(configPath)
	viper.SetConfigFile(filepath.Join(configPath, "appsettings.json"))
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	envSettings := fmt.Sprintf("./configs/appsettings.%s.json", env.EnvName)
	if core.CheckFileExists(envSettings) {
		viper.SetConfigFile(filepath.Join(configPath, fmt.Sprintf("appsettings.%s.json", env.EnvName)))
		err = viper.MergeInConfig()
		if err != nil {
			return nil, err
		}
	}
	err = viper.Unmarshal(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
