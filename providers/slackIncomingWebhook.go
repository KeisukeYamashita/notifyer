package providers

import (
	"errors"

	client "github.com/KeisukeYamashita/notifyer/clients/slack"
	"github.com/KeisukeYamashita/notifyer/config"
)

// SlackIncomingWebhookProviderService ...
type SlackIncomingWebhookProviderService struct {
	Config config.ServiceConfig
}

// GetName ...
func (*SlackIncomingWebhookProviderService) GetName() string {
	return "slackIncomingWebhook"
}

func (ps *SlackIncomingWebhookProviderService) Init() error {
	var err error
	if ps.Config, err = config.NewConfig(ps.GetName()); err != nil {
		return err
	}
	return nil
}

func (ps *SlackIncomingWebhookProviderService) Send(msg string) error {
	conf := &config.SlackIncomingWebhookConfig{}
	switch v := ps.Config.(type) {
	case *config.SlackIncomingWebhookConfig:
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
