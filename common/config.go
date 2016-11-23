package common

type Config struct {
	API struct {
		IP         string `default:""`
		Port       uint   `required:"true" default:"1337"`
		Secret     string `required:"true"`
		Unixsocket string `required:"true"`
	}

	Service struct {
		Unixsocket string `required:"true"`
		Script     []string
	}
}
