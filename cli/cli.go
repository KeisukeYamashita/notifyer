package cli

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/KeisukeYamashita/notifyer/providers"
)

// CLI ...
type CLI struct {
	OutStream, ErrorStream io.Writer
	InputSteam             io.Reader
}

// NewCLI ...
func NewCLI(outStream, errStream io.Writer, inputStream io.Reader) *CLI {
	return &CLI{
		OutStream:   outStream,
		ErrorStream: errStream,
		InputSteam:  inputStream,
	}
}

// Execute ...
func (c *CLI) Execute(args []string) int {
	if len(args) == 0 {
		log.Fatal("please pass command")
	}

	provider := args[0]
	var ok bool
	var providerServicesImportFunc func() providers.ProviderSevice
	if providerServicesImportFunc, ok = providerSevicesImport()[provider]; !ok {
		log.Fatalf("%s service does not exist", provider)
	}
	providerService := providerServicesImportFunc()
	if providerService == nil {
		log.Fatalf("%s provider no supported", provider)
	}
	log.Printf("%s provider selected", provider)

	if err := providerService.Init(); err != nil {
		log.Fatalf("%s init failed: %v", provider, err)
	}
	log.Printf("%s init", provider)

	b, err := ioutil.ReadAll(c.InputSteam)
	if err != nil {
		log.Fatalf("%s faield to read input: %v", provider, err)
	}

	if err := providerService.Send(string(b)); err != nil {
		log.Fatalf("%s sent failed: %v", provider, err)
	}
	log.Printf("%s send request successfully", provider)

	return 0
}

func providerSevicesImport() map[string]func() providers.ProviderSevice {
	list := make(map[string]func() providers.ProviderSevice)

	for _, providerGen := range []func() providers.ProviderSevice{
		newDiscordProvider,
		newSlackProvider,
		newLineBotProvider,
	} {
		list[providerGen().GetName()] = providerGen
	}

	return list
}

func newDiscordProvider() providers.ProviderSevice {
	return &providers.DiscordProviderService{}
}

func newSlackProvider() providers.ProviderSevice {
	return &providers.SlackProviderService{}
}

func newLineBotProvider() providers.ProviderSevice {
	return &providers.LineBotProviderService{}
}
