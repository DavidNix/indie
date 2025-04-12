package asset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	t.Parallel()

	t.Run("no cache key", func(t *testing.T) {
		cacheKey = ""
		imgPath := "example.png"
		want := "example.png"
		got := Path(imgPath)
		require.Equal(t, want, got)
	})

	t.Run("cache key with existing query", func(t *testing.T) {
		cacheKey = "123"
		imgPath := "example.png?size=large"
		want := "example.png?size=large&v=123"
		got := Path(imgPath)
		require.Equal(t, want, got)
	})

	t.Run("cache key without existing query", func(t *testing.T) {
		cacheKey = "123"
		imgPath := "example.png"
		want := "example.png?v=123"
		got := Path(imgPath)
		require.Equal(t, want, got)
	})
}
