package config

type DiscordServiceConfig struct {
	URL string `toml:"url"`
}

func newDiscordConfig() (ServiceConfig, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return conf.Discord, nil
}

func (DiscordServiceConfig) GetName() string {
	return "discord"
}
