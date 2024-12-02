package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mastructure:"DB_DRIVER"`
	DBHost        string `mastructure:"DB_HOST"`
	DBPort        string `mastructure:"DB_PORT"`
	DBUser        string `mastructure:"DB_USER"`
	DBPassword    string `mastructure:"DB_PASSWORD"`
	DBName        string `mastructure:"DB_NAME"`
	WebServerPort string `mastructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mastructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mastructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	jwtSecret, _ := viper.Get("JWT_SECRET").(string)
	cfg.JwtExpiresIn = viper.GetInt("JWT_EXPIRES_IN")
	cfg.TokenAuth = jwtauth.New("HS256", []byte(jwtSecret), nil)
	return cfg, err
}
