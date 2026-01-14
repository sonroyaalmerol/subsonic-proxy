package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) Stream(w http.ResponseWriter, r *http.Request, params opensubsonic.StreamParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostStream(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) Download(w http.ResponseWriter, r *http.Request, params opensubsonic.DownloadParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostDownload(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) Scrobble(w http.ResponseWriter, r *http.Request, params opensubsonic.ScrobbleParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostScrobble(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetTranscodeDecision(w http.ResponseWriter, r *http.Request, params opensubsonic.GetTranscodeDecisionParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetTranscodeStream(w http.ResponseWriter, r *http.Request, params opensubsonic.GetTranscodeStreamParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) HlsM3u8(w http.ResponseWriter, r *http.Request, params opensubsonic.HlsM3u8Params) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostHlsM3u8(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
