package main

import (
	"fmt"
	"os"

	"github.com/Ribas160/gopenvpn/pkg/repository"
	"github.com/Ribas160/gopenvpn/pkg/service"
	"github.com/sirupsen/logrus"
)

const (
	ClientsConfigs       string = "/etc/openvpn/clients/"
	EasyRsa              string = "/etc/openvpn/easy-rsa/"
	ClientsDefaultConfig string = "/etc/openvpn/server/client_default.ovpn"
	ConfigsPrefix        string = "client_"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatalf("Command not specified")

	} else if len(os.Args) < 3 {
		logrus.Fatalf("Client's name is not specified")
	}

	os.Setenv("CLIENTS_CONFIGS", ClientsConfigs)
	os.Setenv("EASY_RSA", EasyRsa)
	os.Setenv("CLIENT_DEFAULT_CONFIG", ClientsDefaultConfig)
	os.Setenv("CONFIGS_PREFIX", ConfigsPrefix)

	args := os.Args

	command := args[1]
	clientName := args[2]

	client := repository.NewClient(repository.Config{
		Name: clientName,
	})

	repos := repository.NewRepository(client)
	services := service.NewServices(repos)

	c := &Command{
		services: services,
	}

	if command == "build" {
		c.buildNew(clientName)

	} else if command == "config" {
		c.createConfig()

	} else {
		logrus.Fatalf("Command " + command + "does not exists")
	}
}

type Command struct {
	services *service.Service
}

func (c *Command) buildNew(clientName string) {
	err := c.services.Client.BuildNew()
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	fmt.Printf("Key was generated: %spki/private/%s.key\n", os.Getenv("EASY_RSA"), clientName)
	fmt.Printf("Certificate was generated: %spki/issued/%s.crt\n", os.Getenv("EASY_RSA"), clientName)

	c.createConfig()
}

func (c *Command) createConfig() {
	config, err := c.services.Client.CreateConfig()
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	fmt.Println("Config was successfully created: " + config)
}
