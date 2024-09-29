package config

type Config struct {
	DBUser    string
	DBPass    string
	DBName    string
	RedisAddr string
	JWTSecret string
}

var AppConfig *Config

func LoadConfig() {
	//viper.SetConfigFile(".env")
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatal("加载失败 .env file")
	//}
	AppConfig = &Config{
		DBUser:    "yuki",
		DBPass:    "ayase",
		DBName:    "geekdemo",
		RedisAddr: "localhost:6379",
		JWTSecret: "c59z",
	}
}
