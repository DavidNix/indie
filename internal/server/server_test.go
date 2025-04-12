package server

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
	os.Exit(m.Run())
}

func ValidCSRFRequest(t *testing.T, b AppBuilder, method, target string, body io.Reader) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	b.Build().ServeHTTP(w, r)
	token := w.Result().Cookies()[0].Value

	req := httptest.NewRequest(method, target, body)
	req.Header.Set("X-CSRF-TOKEN", token)
	for _, cookie := range w.Result().Cookies() {
		req.AddCookie(cookie)
	}
	return req
}

func SaveAndOpenPage(t *testing.T, html string) {
	dir := t.TempDir()
	f, err := os.Create(dir + "/save-and-open-page.html")
	require.NoError(t, err)
	t.Cleanup(func() { f.Close() })
	_, err = io.Copy(f, strings.NewReader(html))
	require.NoError(t, err)

	err = exec.Command("open", f.Name()).Run()
	require.NoError(t, err)

	require.Fail(t, "stopping test to inspect page")
}
