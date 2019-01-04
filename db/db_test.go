package db

import (
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

type item struct {
	Foo string
	Bar string
}

func TestDb(t *testing.T) {
	db, err := NewDb(os.Getenv("REDIS_ADDRESS"), 0)
	require.NoError(t, err)
	require.NotNil(t, db)

	t.Run("int", func(t *testing.T) {
		_, err := db.Set("foo", 1)
		require.NoError(t, err)

		var val int
		_, err = db.Get("foo", &val)
		require.NoError(t, err)

		assert.Equal(t, val, 1)
	})

	t.Run("string", func(t *testing.T) {
		_, err := db.Set("foo", "foo bar baz 1337")
		require.NoError(t, err)

		var val string
		_, err = db.Get("foo", &val)
		require.NoError(t, err)

		assert.Equal(t, val, "foo bar baz 1337")
	})

	t.Run("struct", func(t *testing.T) {
		_, err := db.Set("foo", &item{Foo: "foo", Bar: "baz"})
		require.NoError(t, err)

		var val *item
		_, err = db.Get("foo", &val)
		require.NoError(t, err)

		assert.Equal(t, val, &item{Foo: "foo", Bar: "baz"})
	})

	t.Run("bool", func(t *testing.T) {
		_, err := db.Set("foo", false)
		require.NoError(t, err)

		var val bool
		_, err = db.Get("foo", &val)
		require.NoError(t, err)

		assert.Equal(t, val, false)
	})
}
