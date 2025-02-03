package config

import (
	//godotenv
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type APIConfig struct {
	ApiPort int
}

type TokenConfig struct {
	ApplicationName    string
	JwtSignatureKey    []byte
	JwtSigningMethod   *jwt.SigningMethodHMAC
	AccesTokenLifeTime time.Duration
}

type Config struct {
	DB    DBConfig
	API   APIConfig
	Token TokenConfig
}

func (c *Config) readConfig() error {

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		return err
	}

	c.DB = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}

	c.API = APIConfig{
		ApiPort: apiPort,
	}

	accessTokenLifeTime := time.Duration(1) * time.Hour
	c.Token = TokenConfig{
		ApplicationName:    "Toko Cik Bos",
		JwtSignatureKey:    []byte("TokoCikBosSecretMax"),
		JwtSigningMethod:   jwt.SigningMethodHS256,
		AccesTokenLifeTime: accessTokenLifeTime,
	}

	if c.DB.Host == "" || c.DB.Port == 0 || c.DB.User == "" || c.DB.Password == "" || c.DB.Dbname == "" || c.API.ApiPort == 0 || c.Token.ApplicationName == "" || c.Token.JwtSignatureKey == nil || c.Token.JwtSigningMethod == nil || c.Token.AccesTokenLifeTime == 0 {
		return fmt.Errorf("Required Config")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.readConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
