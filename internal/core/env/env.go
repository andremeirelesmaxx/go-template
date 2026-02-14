package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB     db
	Server server
	App    app
}

func setupEnv() Env {
	app := newApp()
	if app.IsLocal {
		envPath := fmt.Sprintf("%s/.env", app.Root)
		godotenv.Load(envPath)
	}

	return Env{
		DB:     newDb(),
		Server: newServer(),
		App:    app,
	}
}

func NewEnv() Env {
	return setupEnv()
}

func getEnvWithDefault(name string, fallback string) string {
	env := os.Getenv(name)
	if env == "" {
		return fallback
	}

	return env
}
