package configs

import (
	"errors"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	conf *Configuration
	once sync.Once
)

// Config loads configuration using atomic pattern
func Config() *Configuration {
	once.Do(func() {
		conf = load()
	})
	return conf
}

// Configuration ...
type Configuration struct {
	ServerPort         string
	ServerHost         string
	MaxLengthChallenge int
	MinLengthChallenge int
	DifficultyLength   int
	SolutionLength     int
}

func load() *Configuration {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file: ", err)
	}

	var config Configuration

	v := viper.New()
	v.AutomaticEnv()

	config.ServerPort = v.GetString("SERVER_PORT")
	config.ServerHost = v.GetString("SERVER_HOST")
	config.MaxLengthChallenge = v.GetInt("MAX_LENGTH_CHALLENGE")
	config.MinLengthChallenge = v.GetInt("MIN_LENGTH_CHALLENGE")
	config.DifficultyLength = v.GetInt("DIFFICULTY_LENGTH")
	config.SolutionLength = v.GetInt("SOLUTION_LENGTH")

	//validate the configuration
	err = config.validate()
	if err != nil {
		log.Fatal("error validating config: ", err)
	}

	return &config
}

func (c *Configuration) validate() error {
	if c.ServerPort == "" {
		return errors.New("SERVER_PORT required")
	}
	if c.ServerHost == "" {
		return errors.New("SERVER_HOST required")
	}

	// ....

	return nil
}
