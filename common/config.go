package common

import "github.com/jinzhu/configor"

type Deployment struct {
	Name        string   `required:"true"`
	Description string   `default:""`
	Secret      string   `required:"true"`
	User        string   `default:"root"`
	Script      []string `required:"true"`
}

type Config struct {
	API struct {
		IP         string `default:""`
		Port       uint   `default:"1337"`
		Unixsocket string `default:"./service_client.sock"`
	}

	Service struct {
		Unixsocket string `default:"./service.sock"`
	}

	Deployments []Deployment `required:"true"`
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
