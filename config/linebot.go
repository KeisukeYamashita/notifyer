package config

type LineBotServiceConfig struct {
	AccessToken string
	To          string
}

func newLineBotConfig() (ServiceConfig, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return conf.Linebot, nil
}

func (LineBotServiceConfig) GetName() string {
	return "linebot"
}
