package providers

import (
	"errors"

	client "github.com/KeisukeYamashita/notifyer/clients/linebot"
	"github.com/KeisukeYamashita/notifyer/config"
)

// LineBotProviderService ...
type LineBotProviderService struct {
	Config config.ServiceConfig
}

// GetName ...
func (*LineBotProviderService) GetName() string {
	return "linebot"
}

func (ps *LineBotProviderService) Init() error {
	var err error
	if ps.Config, err = config.NewConfig(ps.GetName()); err != nil {
		return err
	}
	return nil
}

// Send ...
func (ps *LineBotProviderService) Send(msg string) error {
	conf := &config.LineBotServiceConfig{}
	switch v := ps.Config.(type) {
	case *config.LineBotServiceConfig:
		conf = v
	default:
		return errors.New("no type")
	}

	client := client.NewClient(conf.AccessToken, conf.To)
	if err := client.Send(msg); err != nil {
		return err
	}

	return nil
}
