package asset

import "strings"

var cacheKey string

// SetCacheKey sets a cache busting key that will be appended to asset URLs.
// A good cache key is your app version or git hash.
// This helps force browsers to download fresh copies of assets when the key changes.
// Warning: Not Goroutine safe. Set once in main().
func SetCacheKey(key string) {
	cacheKey = key
}

// Path appends a cache busting key to the given path URL.
// If a cache key is not set, it returns the original image path unchanged.
func Path(p string) string {
	if cacheKey == "" {
		return p
	}
	if i := strings.Index(p, "?"); i != -1 {
		return p + "&v=" + cacheKey
	}
	return p + "?v=" + cacheKey
}
