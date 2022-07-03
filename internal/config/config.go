package config

import (
	"log"
	"os"

	"github.com/donreno/dashboar/internal/types"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type ConfigurationLoader func() (*types.EnvConfig, *types.Dashboard, error)

func LoadConfiguration() (conf *types.EnvConfig, dashboar *types.Dashboard, err error) {

	log.Println("Initializing Configuration")

	conf, err = loadViperConfig()
	if err != nil {
		log.Printf("Error loading env config %v", err.Error())
		return
	}

	dashboar, err = loadYamlConfigDashboar(conf.DashBoarConfigPath)
	if err != nil {
		log.Printf("Error loading dashboard config %v", err.Error())
		return
	}

	log.Println("Configuration loaded successfully!")

	return
}

func loadViperConfig() (conf *types.EnvConfig, err error) {
	conf = new(types.EnvConfig)

	viper.SetConfigType("env")
	viper.SetDefault("DASHBOAR_CONFIG_FILE", "./samples/config.yaml")
	viper.SetDefault("DASHBOAR_PORT", 8000)
	viper.SetDefault("LIMITER_MAX_REQUESTS", 20)
	viper.SetDefault("LIMITER_EXPIRE_TIME_SECONDS", 5)
	viper.SetDefault("CACHE_EXPIRE_TIME_SECONDS", 60)

	viper.AutomaticEnv()

	viper.Unmarshal(conf)

	return
}

func loadYamlConfigDashboar(path string) (*types.Dashboard, error) {
	dashboar := new(types.Dashboard)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(dashboar)
	if err != nil {
		return nil, err
	}

	return dashboar, nil
}
