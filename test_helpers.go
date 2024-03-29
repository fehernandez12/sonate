package sonate

import "net/http"

// SetURLVars sets the URL variables for the given request, to be accessed via
// sonate.Vars for testing route behaviour. Arguments are not modified, a shallow
// copy is returned.
//
// This API should only be used for testing purposes; it provides a way to
// inject variables into the request context. Alternatively, URL variables
// can be set by making a route that captures the required variables,
// starting a server and sending the request to that server.
func SetURLVars(r *http.Request, val map[string]string) *http.Request {
	return requestWithVars(r, val)
}
