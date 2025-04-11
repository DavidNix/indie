package version

import (
	"cmp"
	"fmt"
	"time"
)

var V string

// Version is expected to be overridden on production builds with ld flags
func Version() string {
	return cmp.Or(V, fmt.Sprintf("dev-%x", time.Now().UnixNano()))
}
