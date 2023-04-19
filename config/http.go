package config

import "github.com/go-fires/example/app/providers/http"

var Http = &http.Config{
	Address: "127.0.0.1:8081",
}
