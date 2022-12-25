package util

import "github.com/spf13/viper"

// Config Stores All configuration of the application.
// The values are read by the viper from a config file or environment variables.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerType    string `mapstructure:"SERVER_TYPE"`
}

// LoadConfig reads configurations from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if config.ServerType == "PRODUCTION" {
		productionDBSource := "postgresql://root:secret@localhost:5432/DCMS?sslmode=disable"
		viper.Set("DB_SOURCE", productionDBSource)
		config.DBSource = productionDBSource
	}
	return
}
