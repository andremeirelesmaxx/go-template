package env

import "os"

type server struct {
	Address string
}

func newServer() server {
	return server{os.Getenv("APP_PORT")}
}
