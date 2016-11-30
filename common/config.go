package common

import "github.com/jinzhu/configor"

type Deployment struct {
	Name        string
	Description string
	Secret      string `required:"true"`
	Script      []string
}

type Config struct {
	API struct {
		IP         string `default:""`
		Port       uint   `required:"true" default:"1337"`
		Secret     string `required:"true"`
		Unixsocket string `required:"true"`
	}

	Service struct {
		Unixsocket string `required:"true"`
	}

	Deployments []Deployment
}

func MakeConfig(path string) *Config {
	var config Config

	// Read config
	configor.Load(&config, path)

	return &config
}

func (config *Config) FindDeploymentByName(name string) *Deployment {
	for _, deployment := range config.Deployments {
		if deployment.Name == name {
			return &deployment
		}
	}

	return nil
}
