package config

type Config struct {
	Mod       string
	RedisHost string
	RedisPort string
}

var GlobalConfig = Config{}

func InitConfig() {

}
