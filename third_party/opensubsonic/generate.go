package opensubsonic

//go:generate npx @redocly/cli bundle https://raw.githubusercontent.com/opensubsonic/open-subsonic-api/refs/heads/main/openapi/openapi.json --output bundled.json --ext json
//go:generate oapi-codegen -config cfg.yaml ./bundled.json
//go:generate rm ./bundled.json
