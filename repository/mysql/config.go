package mysql

import "fmt"

type Config struct {
	Driver string `koanf:"mysql"`
	User   string `koanf:"user"`
	Pass   string `koanf:"password"`
	Host   string `koanf:"host"`
	Port   string `koanf:"port"`
	Name   string `koanf:"name"`
}

func (c Config) BuildURL() string {
	return fmt.Sprintf(`%s:%s@(%s:%s)/%s`, c.User, c.Pass, c.Host, c.Port, c.Name)
}
