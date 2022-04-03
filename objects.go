package delivery

import (
	"net/http"
)

var (
	httpClient *http.Client
)

// http related constants
const (
	listenAddr       = ""
	HTTPReadTimeout  = 60
	HTTPWriteTimeout = 60
)
