package asset

import "strings"

// ImageSrc appends a cache busting key to the given image path URL.
// If a cache key is not set, it returns the original image path unchanged.
// If the image path already contains query parameters, the cache key is appended with &.
// Otherwise, the cache key is appended with ?.
func ImageSrc(imgPath string) string {
	if cacheKey == "" {
		return imgPath
	}
	if i := strings.Index(imgPath, "?"); i != -1 {
		return imgPath + "&v=" + cacheKey
	}
	return imgPath + "?v=" + cacheKey
}
