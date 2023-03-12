// Package realport a demo plugin.
package realport

import (
	"context"
	"net/http"
)

// RealPort a plugin to add remote port to requests.
type RealPort struct {
	next    http.Handler
	headers map[string]string
	name    string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, name string) (http.Handler, error) {

	return &RealPort{
		next: next,
		name: name,
	}, nil
}

func (a *RealPort) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	//realport := strings.Split(req.RemoteAddr, ":")

	req.Header.Set("X-Real-Port", "1212")

	a.next.ServeHTTP(rw, req)
}
