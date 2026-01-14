package handlers

import (
	"net/http"

	"github.com/sonroyaalmerol/subsonic-proxy/internal/api/opensubsonic"
)

func (s *SubsonicProxy) Ping(w http.ResponseWriter, r *http.Request) { s.handleNotImplemented(w, r) }
func (s *SubsonicProxy) PostPing(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) StartScan(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostStartScan(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetScanStatus(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetScanStatus(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) TokenInfo(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostTokenInfo(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetLicense(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetLicense(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) JukeboxControl(w http.ResponseWriter, r *http.Request, params opensubsonic.JukeboxControlParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostJukeboxControl(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetMusicDirectory(w http.ResponseWriter, r *http.Request, params opensubsonic.GetMusicDirectoryParams) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetMusicDirectory(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) GetMusicFolders(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
func (s *SubsonicProxy) PostGetMusicFolders(w http.ResponseWriter, r *http.Request) {
	s.handleNotImplemented(w, r)
}
