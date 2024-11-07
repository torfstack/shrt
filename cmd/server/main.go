package main

import (
	"fmt"
	"github.com/torfstack/shrt/pkg/config"
	"github.com/torfstack/shrt/pkg/http"
	"github.com/torfstack/shrt/pkg/service"
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
