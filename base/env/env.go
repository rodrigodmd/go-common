package env

import "os"

type EnvDefault struct {
	EnvKey       string
	DefaultValue string
}

var BASE_URL_ENV = EnvDefault{
	EnvKey:       "MB_TOOLS_PATH",
	DefaultValue: os.Getenv("HOME") + "/.mb-tool/",
}

func GetEnv(env EnvDefault) string {
	envValue := os.Getenv(env.EnvKey)
	if envValue != "" {
		return envValue
	}
	return env.DefaultValue
}

func BaseUrl() string {
	return GetEnv(BASE_URL_ENV)
}
