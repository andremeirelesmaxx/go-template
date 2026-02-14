package env

import "os"

type db struct {
	User     string
	Password string
	Name     string
	Port     string
	Host     string
}

func newDb() db {
	return db{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
	}
}
