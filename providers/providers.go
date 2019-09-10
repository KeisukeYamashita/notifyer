package providers

type ProviderSevice interface {
	GetName() string
	Init() error
	Send(msg string) error
}
