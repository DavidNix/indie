package main

import (
	"fmt"

	"github.com/DavidNix/indie/internal/version"
)

func main() {
	fmt.Println("Version: %s", version.Version)
}
