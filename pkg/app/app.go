package app

import "go1/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}
