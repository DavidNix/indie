package version

import (
	"fmt"
	"time"
)

// Version is expected to be overridden on production builds with ld flags
var Version = fmt.Sprintf("dev-%x", time.Now().UnixNano())
