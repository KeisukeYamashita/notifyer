package providers

import (
	"errors"

	client "github.com/KeisukeYamashita/notifyer/clients/slack"
	"github.com/KeisukeYamashita/notifyer/config"
)

// SlackIncomingWebhookProviderService ...
type SlackProviderService struct {
	Config config.ServiceConfig
}

// GetName ...
func (*SlackProviderService) GetName() string {
	return "slack"
}

func (ps *SlackProviderService) Init() error {
	var err error
	if ps.Config, err = config.NewConfig(ps.GetName()); err != nil {
		return err
	}
	return nil
}

func (ps *SlackProviderService) Send(msg string) error {
	conf := &config.SlackConfig{}
	switch v := ps.Config.(type) {
	case *config.SlackConfig:
		conf = v
	default:
		return errors.New("no type")
	}

	client := client.NewClient(conf.URL)
	if err := client.Send(msg); err != nil {
		return err
	}

	return nil
}
