package realport

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRealPort(t *testing.T) {

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, "realport")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req)
}

func assertHeader(t *testing.T, req *http.Request) {
	t.Helper()

	if req.Header.Get("X-Real-Port") == "" {
		t.Errorf("invalid header value: %s", req.Header.Get("X-Real-Port"))
	}
}
