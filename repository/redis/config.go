package redis

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}
