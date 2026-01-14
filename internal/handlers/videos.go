package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) GetVideoInfo(w http.ResponseWriter, r *http.Request, params opensubsonic.GetVideoInfoParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetVideoInfo(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetVideos(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetVideos(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetCaptions(w http.ResponseWriter, r *http.Request, params opensubsonic.GetCaptionsParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetCaptions(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
