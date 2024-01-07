package config

import "github.com/spf13/viper"

type ConfigData struct {
	WebServerPort            string `mapstructure:"WEB_SERVER_PORT"`
	RedisHost                string `mapstructure:"REDIS_HOST"`
	RedisPort                string `mapstructure:"REDIS_PORT"`
	RedisDB                  int    `mapstructure:"REDIS_DB"`
	RedisPassword            string `mapstructure:"REDIS_PASSWORD"`
	RedisRateLimitTTL        int    `mapstructure:"REDIS_RATE_LIMIT_TTL"`
	MaxRequestIPSecond       int    `mapstructure:"MAX_REQUESTS_PER_IP_PER_SECOND"`
	MaxRequestTokenSecond    int    `mapstructure:"MAX_REQUESTS_PER_TOKEN_PER_SECOND"`
	RequestBlockingTimeIP    int    `mapstructure:"REQUEST_BLOCKING_TIME_IP"`
	RequestBlockingTimeToken int    `mapstructure:"REQUEST_BLOCKING_TIME_TOKEN"`
}

func LoadConfig(path string) (*ConfigData, error) {
	var configData *ConfigData
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return configData, err
	}

	err = viper.Unmarshal(&configData)
	if err != nil {
		return configData, err
	}

	return configData, nil
}
