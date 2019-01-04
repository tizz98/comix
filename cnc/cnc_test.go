package cnc

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tizz98/comix/db"
)

func TestService_Ping(t *testing.T) {
	database, err := db.NewDb(os.Getenv("REDIS_ADDRESS"), 0)
	require.NoError(t, err)
	require.NotNil(t, database)

	uri, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	service := &Service{db: database, distUrl: uri}

	t.Run("NotOk", func(t *testing.T) {
		_, err := service.Ping(nil, &PingMsg{Ok: false, StatusMessage: "unable to apply update", ClientId: "123"})
		require.NoError(t, err)

		status, err := service.getClientStatus("123")
		require.NoError(t, err)

		assert.Equal(t, "unable to apply update", status)
	})

	t.Run("NoUpdate", func(t *testing.T) {
		err := service.setClientCurrentVersion("1234", 3)
		require.NoError(t, err)

		err = service.setLatestFileVersion(3)
		require.NoError(t, err)

		resp, err := service.Ping(nil, &PingMsg{Ok: true, ClientId: "1234"})
		require.NoError(t, err)

		assert.False(t, resp.HasUpdate)
	})

	t.Run("UpdateNeeded", func(t *testing.T) {
		err := service.setClientCurrentVersion("12345", 2)
		require.NoError(t, err)

		err = service.setLatestFileVersion(3)
		require.NoError(t, err)

		// mock function
		getChecksum = func(url string) ([]byte, error) {
			return []byte("123"), nil
		}

		resp, err := service.Ping(nil, &PingMsg{Ok: true, ClientId: "12345"})
		require.NoError(t, err)

		assert.True(t, resp.HasUpdate)
		assert.Equal(t, "https://example.com/comix-pi/3", resp.Url)
		assert.Equal(t, []byte("123"), resp.Checksum)
	})
}
