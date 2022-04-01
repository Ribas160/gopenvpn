package main

import (
	"fmt"
	"os"

	"github.com/Ribas160/gopenvpn/pkg/repository"
	"github.com/Ribas160/gopenvpn/pkg/service"
	"github.com/sirupsen/logrus"
)

const (
	ClientsConfigs       string = "/Applications/XAMPP/xamppfiles/htdocs/projects/myProjects/gopenvpn/clients/"
	EasyRsa              string = "/Applications/XAMPP/xamppfiles/htdocs/projects/myProjects/gopenvpn/easy-rsa/"
	ClientsDefaultConfig string = "/Applications/XAMPP/xamppfiles/htdocs/projects/myProjects/gopenvpn/server/client_default.ovpn"
	ConfigsPrefix        string = "client_"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatalf("Client's name is not specified")
	}

	os.Setenv("CLIENTS_CONFIGS", ClientsConfigs)
	os.Setenv("EASY_RSA", EasyRsa)
	os.Setenv("CLIENT_DEFAULT_CONFIG", ClientsDefaultConfig)
	os.Setenv("CONFIGS_PREFIX", ConfigsPrefix)

	args := os.Args

	clientName := args[1]

	client := repository.NewClient(repository.Config{
		Name: clientName,
	})

	repos := repository.NewRepository(client)
	services := service.NewServices(repos)

	config, err := services.ClientConfig.Create()
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	fmt.Println("Config was successfully created: " + config)
}
