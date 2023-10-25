package server

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	t.Parallel()

	app := NewApp()
	r := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(r)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
}
