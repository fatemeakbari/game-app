package cfg

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load() *Config {

	var k = koanf.New(".")

	k.Load(confmap.Provider(map[string]interface{}{
		"auth.token_secret_key":                "dsfkashfkhsdfkshfasfjsflflsf",
		"http_server.server_shutdown_duration": "5s",
	}, "."), nil)

	//read from config.yml
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		panic(err)
	}

	//TODO not working
	//merge from environment variables
	//k.Load(env.Provider("GANEAPP_", ".", func(s string) string {
	//	return strings.Replace(strings.ToLower(
	//		strings.TrimPrefix(s, "GAMEAPP_")), "_", ".", -1)
	//}), nil)

	var config Config

	//TODO unmarshal duration
	err := k.Unmarshal("", &config)

	if err != nil {
		panic(err.Error())
	}

	return &config

}
