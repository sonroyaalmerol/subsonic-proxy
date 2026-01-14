package opensubsonic

//go:generate go tool oapi-codegen -config cfg.yaml ./openapi/openapi.json
//go:generate sed -i "/GetTranscodeDecision/! s/200JSONResponse/200JSONResponse_SubsonicResponse/g" ../../internal/api/opensubsonic/opensubsonic.go
//go:generate sed -i "s/_SubsonicResponse_SubsonicResponse/_SubsonicResponse/g" ../../internal/api/opensubsonic/opensubsonic.go
//go:generate sed -i "s/GetTranscodeDecision200JSONResponseTranscodeDecisionSourceStreamProtocol/string/g" ../../internal/api/opensubsonic/opensubsonic.go
//go:generate sed -i "s/GetTranscodeDecision200JSONResponseTranscodeDecisionTranscodeStreamProtocol/string/g" ../../internal/api/opensubsonic/opensubsonic.go
