package assets

import (
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(root fs.FS) echo.HandlerFunc {
	static := http.StripPrefix("/", http.FileServerFS(root))
	return echo.WrapHandler(static)
}
