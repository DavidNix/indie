package database

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDB_Migrate(t *testing.T) {
	t.Parallel()

	dbPath := filepath.Join(t.TempDir(), "test.db")
	db, err := Open(dbPath)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	err = db.Migrate()
	require.NoError(t, err)

	err = db.Migrate()
	require.NoError(t, err, "migrations should be idempotent")
}
