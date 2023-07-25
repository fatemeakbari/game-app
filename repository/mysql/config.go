package mysql

import "fmt"

type Config struct {
	Driver string
	User   string
	Pass   string
	Host   string
	Port   string
	Name   string
}

func (c Config) buildURL() string {
	return fmt.Sprintf(`%s:%s@(%s:%s)/%s`, c.User, c.Pass, c.Host, c.Port, c.Name)
}
