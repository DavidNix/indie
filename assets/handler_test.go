package assets

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

//go:embed testdata
var embeddedFS embed.FS

//go:embed testdata/image.png
var imageFixture []byte

func TestHandler(t *testing.T) {
	t.Parallel()

	const testFileName = "image.png"

	t.Run("with embedded fs", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/"+testFileName, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		sub, err := fs.Sub(embeddedFS, "testdata")
		require.NoError(t, err)

		handler := Handler(sub)
		err = handler(c)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, rec.Code)
		respBody, err := io.ReadAll(rec.Body)
		require.NoError(t, err)
		require.Equal(t, imageFixture, respBody)
	})

	t.Run("with os fs", func(t *testing.T) {
		dirFS := os.DirFS("testdata")

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/"+testFileName, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := Handler(dirFS)
		err := handler(c)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, rec.Code)
		respBody, err := io.ReadAll(rec.Body)
		require.NoError(t, err)
		require.Equal(t, imageFixture, respBody)
	})

	t.Run("not found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/testdata/doesnotexist.txt", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := Handler(embeddedFS)
		err := handler(c)
		require.NoError(t, err)

		require.Equal(t, http.StatusNotFound, rec.Code)
	})
}
