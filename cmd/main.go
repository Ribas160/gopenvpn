package main

import (
	"fmt"
	"os"

	"github.com/Ribas160/gopenvpn/pkg/repository"
	"github.com/Ribas160/gopenvpn/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatalf("Client's name is not specified")
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

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
