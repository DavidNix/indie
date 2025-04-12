package asset

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler returns an echo.HandlerFunc that serves static files from the given fs.FS.
// The stripPrefix parameter is removed from the beginning of requested URLs before looking up files.
func Handler(stripPrefix string, root fs.FS) echo.HandlerFunc {
	h := http.StripPrefix(stripPrefix, http.FileServerFS(root))
	return echo.WrapHandler(h)
}
