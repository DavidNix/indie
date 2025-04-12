//go:build !prod

package server

import (
	"io/fs"
	"os"
)

func assetsFS() fs.FS {
	return os.DirFS("internal/server/static")
}
