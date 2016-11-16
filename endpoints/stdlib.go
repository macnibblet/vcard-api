package endpoints

import (
	"encoding/json"
	"net/http"
)

var (
	RESP_INTERNAL_SERVER_ERROR, _ = json.Marshal(map[string]string{
		"error": "An internal error occurred, please try again later",
	})

	RESP_MALFORMED_REQUEST, _ = json.Marshal(map[string]string{
		"error": "The request was malformed, get out of my fucking api",
	})
)

type WrappedHandlerFunc struct {
	handleFunc http.HandlerFunc
}

func (f *WrappedHandlerFunc) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	f.handleFunc(resp, req)
}

func wrapHandlerFunc(handleFunc http.HandlerFunc) http.Handler {
	return &WrappedHandlerFunc{
		handleFunc: handleFunc,
	}
}
