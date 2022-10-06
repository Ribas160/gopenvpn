package repository

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
)

type ClientReposiory struct {
	name string
}

func newClientRepository(c *Config) *ClientReposiory {
	return &ClientReposiory{
		name: c.Name,
	}
}

func (r *ClientReposiory) CreateConfig() (string, error) {
	err := os.MkdirAll(os.Getenv("CLIENTS_CONFIGS"), os.ModePerm)
	if err != nil {
		return "", err
	}

	config, err := newFilesRepository().Read(os.Getenv("CLIENT_DEFAULT_CONFIG"))
	if err != nil {
		return "", err
	}

	cert, err := newFilesRepository().Read(os.Getenv("EASY_RSA") + "pki/issued/" + r.name + ".crt")
	if err != nil {
		return "", err
	}
	regex := regexp.MustCompile(`(?s).*?-----BEGIN CERTIFICATE-----`)
	cert = regex.ReplaceAllString(cert, "-----BEGIN CERTIFICATE-----")

	key, err := newFilesRepository().Read(os.Getenv("EASY_RSA") + "pki/private/" + r.name + ".key")
	if err != nil {
		return "", err
	}

	ca, err := newFilesRepository().Read(os.Getenv("EASY_RSA") + "pki/ca.crt")
	if err != nil {
		return "", err
	}

	regexCa := regexp.MustCompile("(?s)<ca>.*?</ca>")
	config = regexCa.ReplaceAllString(config, "<ca>\n"+ca+"</ca>")

	regexKey := regexp.MustCompile("(?s)<key>.*?</key>")
	config = regexKey.ReplaceAllString(config, "<key>\n"+key+"</key>")

	regexCert := regexp.MustCompile("(?s)<cert>.*?</cert>")
	config = regexCert.ReplaceAllString(config, "<cert>\n"+cert+"</cert>")

	f, err := os.OpenFile(os.Getenv("CLIENTS_CONFIGS")+os.Getenv("CONFIGS_PREFIX")+r.name+".ovpn", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}

	defer f.Close()

	f.WriteString(config)

	return f.Name(), nil
}

func (r *ClientReposiory) BuildNew() error {
	cmd := exec.Command(os.Getenv("EASY_RSA")+"easyrsa", "build-client-full", r.name, "nopass")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
