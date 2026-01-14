package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) GetLyrics(w http.ResponseWriter, r *http.Request, params opensubsonic.GetLyricsParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetLyrics(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetLyricsBySongId(w http.ResponseWriter, r *http.Request, params opensubsonic.GetLyricsBySongIdParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetLyricsBySongId(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
