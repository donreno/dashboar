package config

import (
	"log"

	"github.com/donreno/dashboar/internal/types"
	"github.com/spf13/viper"
)

type ConfigurationLoader func() (*types.EnvConfig, *types.Dashboard, error)

func LoadConfiguration() (conf *types.EnvConfig, dashboar *types.Dashboard, err error) {

	log.Println("Initializing Configuration")

	conf, err = loadViperConfig()
	if err != nil {
		log.Printf("Error loading env config %v", err.Error())
		return
	}

	dashboar, err = loadMockDashboard()
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

func loadMockDashboard() (*types.Dashboard, error) {
	dashboard := &types.Dashboard{
		Title: "My beautiful dashboard",
		Categories: []types.Category{
			{
				Name: "Networking",
				Icon: "bi-hdd-network",
				Entries: []types.Entry{
					{
						Name:      "PI-hole",
						Icon:      "bi-ethernet",
						Url:       "http://192.168.4.23:9000",
						UseNewTab: true,
					},
					{
						Name:      "NAS",
						Icon:      "bi-hdd-stack",
						Url:       "http://192.168.4.23:5000",
						UseNewTab: true,
					},
				},
			},
			{
				Name: "Coding",
				Icon: "bi-code-slash",
				Entries: []types.Entry{
					{
						Name:      "Github",
						Icon:      "bi-file-code",
						Url:       "http://github.com/donreno",
						UseNewTab: true,
					},
				},
			},
			{
				Name: "Networking",
				Icon: "bi-hdd-network",
				Entries: []types.Entry{
					{
						Name:      "PI-hole",
						Icon:      "bi-ethernet",
						Url:       "http://192.168.4.23:9000",
						UseNewTab: true,
					},
					{
						Name:      "NAS",
						Icon:      "bi-hdd-stack",
						Url:       "http://192.168.4.23:5000",
						UseNewTab: true,
					},
					{
						Name:      "Github",
						Icon:      "bi-file-code",
						Url:       "http://github.com/donreno",
						UseNewTab: true,
					},
				},
			},
			{
				Name: "Coding 2",
				Icon: "bi-code-slash",
				Entries: []types.Entry{
					{
						Name:      "Github",
						Icon:      "bi-file-code",
						Url:       "http://github.com/donreno",
						UseNewTab: true,
					},
				},
			},
		},
	}

	return dashboard, nil
}
