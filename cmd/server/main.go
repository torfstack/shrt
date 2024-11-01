package main

import (
	"fmt"
	"shrt/pkg/config"
	"shrt/pkg/http"
	"shrt/pkg/service"
)

func main() {
	cfg := config.ParseConfig()
	if localhostBaseUrl != "" {
		fmt.Println("WARNING: Using localhost base url")
		cfg.BaseUrl = localhostBaseUrl
	}
	svc := service.NewShrtService(cfg)
	http.StartServer(svc)
}

// -ldflags "-X main.localhostBaseUrl=http://localhost:8080"
var localhostBaseUrl string
