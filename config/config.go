package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var configurations *Config

type DBConfig struct {
	Host string
	Port int
	Name string
	User string
	Password string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB &DBConfig
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION is not set")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME is not set")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP_PORT is not set")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT_SECRET_KEY is not set")
		os.Exit(1)
	}

	host := os.Getenv("HOST")
	if host == "" {
		fmt.Println("host is not set")
		os.Exit(1)
	}

	dbport := os.Getenv("PORT")
	if dbport == "" {
		fmt.Println("DB port is not set")
		os.Exit(1)
	}

	dbprt, err := strconv.ParseInt(dbport, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	name := os.Getenv("NAME")
	if name == "" {
		fmt.Println("name is not set")
		os.Exit(1)
	}

	user := os.Getenv("USER")
	if user == "" {
		fmt.Println("user is not set")
		os.Exit(1)
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		fmt.Println("password is not set")
		os.Exit(1)
	}

	enable_ssl_mode := os.Getenv("ENABLE_SSL_MODE")

	enblSSLMode, err := strconv.ParesBool(enable_ssl_mode)
if err != nil {
	fmt.Println("Invalid enable ssl mode value")
	os.Exit(1)
}


	if err != nil {
		fmt.Println("HTTP_PORT is not a valid number")
		os.Exit(1)
	}

	dbConfig := &dbConfig{
		Host: Host,
		Port: int(dbprt),
		Name: Name,
		User: User,
		Password: Password,
		EnableSSLMODE: enblSSLMode
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB: dbConfig
	}
}

func GetConfig() *Config {

	if configurations == nil {
		loadConfig()
	}
	return configurations
}
