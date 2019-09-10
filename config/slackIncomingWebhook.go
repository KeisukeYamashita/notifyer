package config

type SlackIncomingWebhookConfig struct {
	URL string `toml:"url"`
}

func newSlackIncomingWebhookConfig() (ServiceConfig, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return conf.SlackIncomingWebhook, nil
}

func (SlackIncomingWebhookConfig) GetName() string {
	return "slackIncomingWebhook"
}
