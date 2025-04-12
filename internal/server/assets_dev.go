//go:build !prod

package server

import (
	"io/fs"
	"os"
)

func assetsFS() fs.FS {
	dir := os.DirFS("internal/server/public")
	sub, err := fs.Sub(dir, "internal/server/public")
	if err != nil {
		panic(err)
	}
	return sub
}
