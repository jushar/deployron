package config

type Config struct {
	Api struct {
		IP     string `default:""`
		Port   uint   `required:"true" default:"1337"`
		Secret string `required:"true"`
	}

	Service struct {
		Unixsocket string `required:"true"`
		Script     string `default:"./deploy.sh"`
	}
}
