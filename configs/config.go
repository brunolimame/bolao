package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

func NewConfig() *conf {
	return cfg
}

type conf struct {
	dbDriver      string
	dbHost        string
	dbPort        string
	dbUser        string
	dbPassword    string
	dbName        string
	webServerPort int
	jWTSecret     string
	jWTExpiresIn  int
	tokenAuth     *jwtauth.JWTAuth
}

func init() {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(true)
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg.dbDriver = viper.GetString("DB_DRIVER")
	cfg.dbHost = viper.GetString("DB_HOST")
	cfg.dbPort = viper.GetString("DB_PORT")
	cfg.dbUser = viper.GetString("DB_USER")
	cfg.dbPassword = viper.GetString("DB_PASSWORD")
	cfg.dbName = viper.GetString("DB_NAME")
	cfg.webServerPort = viper.GetInt("WEB_SERVER_PORT")
	cfg.jWTSecret = viper.GetString("JWT_SECRET")
	cfg.jWTExpiresIn = viper.GetInt("JET_EXPIRES_IN")
	cfg.tokenAuth = jwtauth.New("HS256", []byte(cfg.jWTSecret), nil)

}

func (c *conf) GetDBDriver() string {
	return c.dbDriver
}
func (c *conf) GetDBHost() string {
	return c.dbHost
}
func (c *conf) GetDBPort() string {
	return c.dbPort
}
func (c *conf) GetDBUser() string {
	return c.dbUser
}
func (c *conf) GetDBPassword() string {
	return c.dbPassword
}
func (c *conf) GetDBName() string {
	return c.dbName
}
func (c *conf) GetWebServerPort() int {
	return c.webServerPort
}
func (c *conf) GetJwtSecret() string {
	return c.jWTSecret
}
func (c *conf) GetJwtExpiresIn() int {
	return c.jWTExpiresIn
}
