package handlers

import (
	"fmt"
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

// TODO: implement strict server interface instead
var _ opensubsonic.ServerInterface = (*SubsonicProxy)(nil)

type SubsonicProxy struct{}

func (s *SubsonicProxy) handleNotImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint hit: %s %s\n", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"subsonic-response": {"status": "ok", "version": "1.16.1"}}`))
}

func NewSubsonicProxy() *SubsonicProxy {
	return &SubsonicProxy{}
}

