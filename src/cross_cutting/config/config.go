package config

import (
	"errors"
	"os"

	"github.com/subosito/gotenv"
)

var (
	Server struct {
		Host string
		Port string
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
)

func Load() (err error) {
	err = gotenv.Load()

	if err != nil {
		return err
	}

	if err = loadServer(); err != nil {
		return err
	}

	if err = loadDB(); err != nil {
		return err
	}

	return nil
}

func loadServer() (err error) {
	if value, found := os.LookupEnv("SERVER_HOST"); found {
		Server.Host = value
	} else {
		return errors.New("environment variable 'SERVER_HOST' is required")
	}

	if value, found := os.LookupEnv("SERVER_PORT"); found {
		Server.Port = value
	} else {
		return errors.New("environment variable 'SERVER_PORT' is required")
	}

	return nil
}

func loadDB() (err error) {
	if value, found := os.LookupEnv("DATABASE_HOST"); found {
		DB.Host = value
	} else {
		return errors.New("environment variable 'DATABASE_HOST' is required")
	}

	if value, found := os.LookupEnv("DATABASE_PORT"); found {
		DB.Port = value
	} else {
		return errors.New("environment variable 'DATABASE_PORT' is required")
	}

	if value, found := os.LookupEnv("DATABASE_USER"); found {
		DB.User = value
	} else {
		return errors.New("environment variable 'DATABASE_USER' is required")
	}

	if value, found := os.LookupEnv("DATABASE_PASSWORD"); found {
		DB.Password = value
	} else {
		return errors.New("environment variable 'DATABASE_PASSWORD' is required")
	}

	if value, found := os.LookupEnv("DATABASE_NAME"); found {
		DB.DBName = value
	} else {
		return errors.New("environment variable 'DATABASE_NAME' is required")
	}

	return nil
}
