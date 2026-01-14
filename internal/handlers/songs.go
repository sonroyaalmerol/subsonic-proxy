package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) GetSong(w http.ResponseWriter, r *http.Request, params opensubsonic.GetSongParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetSong(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetSongsByGenre(w http.ResponseWriter, r *http.Request, params opensubsonic.GetSongsByGenreParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetSongsByGenre(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetCoverArt(w http.ResponseWriter, r *http.Request, params opensubsonic.GetCoverArtParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetCoverArt(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetGenres(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetGenres(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
