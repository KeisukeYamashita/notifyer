package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

var (
	defaultNotifyerPathEnvVar = "NOTIFYER_PATH"
	defaultNotifyerPath       = os.Getenv("HOME") + "/.notifyer.toml"
)

type TomlConfig struct {
	Default              string
	Discord              *DiscordServiceConfig
	SlackIncomingWebhook *SlackIncomingWebhookConfig
	Linebot              *LineBotServiceConfig
}

type ServiceConfig interface {
	GetName() string
}

func NewConfig(provider string) (ServiceConfig, error) {
	config, err := importConfig()[provider]()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func importConfig() map[string]func() (ServiceConfig, error) {
	configs := make(map[string]func() (ServiceConfig, error))

	// TODO(KeisukeYamahita): Use GetName() for keys
	for name, config := range map[string]func() (ServiceConfig, error){
		"discord":              newDiscordConfig,
		"slackIncomingWebhook": newSlackIncomingWebhookConfig,
		"linebot":              newLineBotConfig,
	} {
		configs[name] = config
	}

	return configs
}

func loadConfig() (*TomlConfig, error) {
	var conf TomlConfig
	loadPath := os.Getenv(defaultNotifyerPathEnvVar)
	if loadPath == "" {
		loadPath = defaultNotifyerPath
	}

	if _, err := toml.DecodeFile(loadPath, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
