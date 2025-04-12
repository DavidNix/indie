//go:build prod

package server

import (
	"embed"

	"io/fs"
)

//go:embed static
var staticFS embed.FS

func assetsFS() fs.FS {
	sub, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	return sub
}
