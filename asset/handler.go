package asset

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(stripPrefix string, root fs.FS) echo.HandlerFunc {
	h := http.StripPrefix(stripPrefix, http.FileServerFS(root))
	return echo.WrapHandler(h)
}
