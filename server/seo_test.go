package server

import (
	"fmt"
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
	require.Contains(t, w.Body.String(), "User-agent: *")
	require.Contains(t, w.Body.String(), fmt.Sprintf("Sitemap: %s/sitemap.xml", baseURL))
}

func TestSitemap(t *testing.T) {
	t.Parallel()

	app := NewApp(nil)

	req := httptest.NewRequest(http.MethodGet, "/sitemap.xml", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
	require.Equal(t, "application/xml; charset=UTF-8", w.Header().Get("Content-Type"))
}
