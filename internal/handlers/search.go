package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) Search(w http.ResponseWriter, r *http.Request, params opensubsonic.SearchParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostSearch(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) Search2(w http.ResponseWriter, r *http.Request, params opensubsonic.Search2Params) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostSearch2(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) Search3(w http.ResponseWriter, r *http.Request, params opensubsonic.Search3Params) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostSearch3(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
