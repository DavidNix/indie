package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DavidNix/indie/ent"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func setupDatabase(t require.TestingT) *ent.Client {
	const (
		driver = "sqlite3"
		dbURL  = "file:indie?mode=memory&cache=shared&_fk=1"
	)
	client, err := ent.Open(driver, dbURL)
	require.NoError(t, err)

	err = client.Schema.Create(context.Background())
	require.NoError(t, err)

	return client
}

func TestExample(t *testing.T) {
	t.Parallel()

	client := setupDatabase(t)

	app := NewApp(client)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
}
