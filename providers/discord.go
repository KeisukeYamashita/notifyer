package providers

import (
	"errors"

	client "github.com/KeisukeYamashita/notifyer/clients/discord"
	"github.com/KeisukeYamashita/notifyer/config"
)

const (
	SupportCodeBlock = true
)

// DiscordProviderService ...
type DiscordProviderService struct {
	Config config.ServiceConfig
}

// GetName ...
func (*DiscordProviderService) GetName() string {
	return "discord"
}

// Init ...
func (ps *DiscordProviderService) Init() error {
	var err error
	if ps.Config, err = config.NewConfig(ps.GetName()); err != nil {
		return err
	}
	return nil
}

// Send ...
func (ps *DiscordProviderService) Send(msg string) error {
	conf := &config.DiscordServiceConfig{}
	switch v := ps.Config.(type) {
	case *config.DiscordServiceConfig:
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
