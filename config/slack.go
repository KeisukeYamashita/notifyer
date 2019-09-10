package config

type SlackConfig struct {
	URL string `toml:"url"`
}

func newSlackConfig() (ServiceConfig, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return conf.Slack, nil
}

func (SlackConfig) GetName() string {
	return "slackIncomingWebhook"
}
