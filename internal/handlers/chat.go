package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) GetChatMessages(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetChatMessages(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetAddChatMessage(w http.ResponseWriter, r *http.Request, params opensubsonic.GetAddChatMessageParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostAddChatMessage(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
