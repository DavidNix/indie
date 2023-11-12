package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRobotsTxt(t *testing.T) {
	t.Parallel()

	app := NewApp(nil)

	req := httptest.NewRequest(http.MethodGet, "/robots.txt", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
	require.NotEmpty(t, w.Body.String())
}
