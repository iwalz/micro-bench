package config

import "github.com/micro/cli"

type Config struct {
	Count    int
	User     string
	Password string
	Database string
	Endpoint string
}

func FromContext(c *cli.Context) Config {
	conf := Config{}
	conf.Count = c.Int("count")
	conf.User = c.String("user")
	conf.Password = c.String("password")
	conf.Database = c.String("database")
	conf.Endpoint = c.String("endpoint")
	return conf
}
