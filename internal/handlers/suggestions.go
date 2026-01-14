package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) GetRandomSongs(w http.ResponseWriter, r *http.Request, params opensubsonic.GetRandomSongsParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetRandomSongs(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetSimilarSongs(w http.ResponseWriter, r *http.Request, params opensubsonic.GetSimilarSongsParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetSimilarSongs(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetSimilarSongs2(w http.ResponseWriter, r *http.Request, params opensubsonic.GetSimilarSongs2Params) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetSimilarSongs2(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
