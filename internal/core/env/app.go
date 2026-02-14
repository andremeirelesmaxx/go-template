package env

type app struct {
	IsLocal bool
	Root    string
}

func newApp() app {
	isLocal := getEnvWithDefault("APP_IS_LOCAL", "true") == "true"

	return app{
		Root:    getEnvWithDefault("APP_ROOT", "/app"),
		IsLocal: isLocal,
	}
}
