package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
	"github.com/sonroyaalmerol/subsonic-proxy/internal/config"
	"github.com/sonroyaalmerol/subsonic-proxy/internal/handlers"
)

func main() {
	server := handlers.NewSubsonicProxy()

	r := http.NewServeMux()

	h := opensubsonic.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf(":%d", config.AppConfig.Server.Port),
	}

	log.Fatal(s.ListenAndServe())
}
