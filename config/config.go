package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type InfoType int

const (
	InfoTypeRSS InfoType = iota
	InfoTypeICal
	InfoPretalx
	InfoHubAssemblies
	InfoHubEvents
)

type Info struct {
	Name string
	URL  string
	Type InfoType
}

type Event struct {
	Name  string
	Infos []Info
}

type Server struct {
	GopherDir  string
	GopherPort int
	SearchPort int
	Hostname   string
}

type Config struct {
	Server Server
	Events []Event
}

func LoadConfig(filepath string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filepath, &config); err != nil {
		return Config{}, errors.New("Failed to load config file: " + err.Error())
	}
	return config, nil
}

func (c *Config) SaveConfig(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return errors.New("Failed to create config file: " + err.Error())
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(c)
	if err != nil {
		return errors.New("Failed to encode config file: " + err.Error())
	}
	return nil
}

type Secrets map[string]string

func LoadSecrets(filepath string) (Secrets, error) {
	var secrets Secrets
	if _, err := toml.DecodeFile(filepath, &secrets); err != nil {
		return Secrets{}, errors.New("Failed to load secrets file: " + err.Error())
	}
	return secrets, nil
}

func (s *Secrets) SaveSecrets(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return errors.New("Failed to create secrets file: " + err.Error())
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(s)
	if err != nil {
		return errors.New("Failed to encode secrets file: " + err.Error())
	}
	return nil
}
