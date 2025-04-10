package asset

var cacheKey string

// SetCacheKey sets a cache busting key that will be appended to asset URLs.
// This helps force browsers to download fresh copies of assets when the key changes.
func SetCacheKey(key string) {
	cacheKey = key
}
