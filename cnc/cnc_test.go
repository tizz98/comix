package cnc

import (
	"net/url"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tizz98/comix/db"
)

func TestService_Ping(t *testing.T) {
	database, err := db.NewDb(os.Getenv("REDIS_ADDRESS"), 0)
	require.NoError(t, err)
	require.NotNil(t, database)

	uri, err := url.Parse("https://example.com/foo/")
	require.NoError(t, err)

	service := &Service{db: database, distUrl: uri}

	t.Run("NotOk", func(t *testing.T) {
		_, err := service.Ping(nil, &PingMsg{Ok: false, StatusMessage: "unable to apply update", ClientId: "123"})
		require.NoError(t, err)

		status, err := service.getClientStatus("123")
		require.NoError(t, err)

		assert.Equal(t, "unable to apply update", status.Message)
	})

	t.Run("NoUpdate", func(t *testing.T) {
		err := service.setClientCurrentVersion("1234", 3)
		require.NoError(t, err)

		err = service.SetLatestFileVersion(3)
		require.NoError(t, err)

		resp, err := service.Ping(nil, &PingMsg{Ok: true, ClientId: "1234"})
		require.NoError(t, err)

		assert.False(t, resp.HasUpdate)
	})

	t.Run("UpdateNeeded", func(t *testing.T) {
		err := service.setClientCurrentVersion("12345", 2)
		require.NoError(t, err)

		err = service.SetLatestFileVersion(3)
		require.NoError(t, err)

		// mock function
		getChecksum = func(url string) ([]byte, error) {
			return []byte("123"), nil
		}

		resp, err := service.Ping(nil, &PingMsg{Ok: true, ClientId: "12345"})
		require.NoError(t, err)

		assert.True(t, resp.HasUpdate)
		assert.Equal(t, "https://example.com/foo/comix-pi/3", resp.Url)
		assert.Equal(t, []byte("123"), resp.Checksum)
	})
}

func TestService_GetClients(t *testing.T) {
	database, err := db.NewDb(os.Getenv("REDIS_ADDRESS"), 0)
	require.NoError(t, err)
	require.NotNil(t, database)

	uri, err := url.Parse("https://example.com/foo/")
	require.NoError(t, err)

	service := &Service{db: database, distUrl: uri}

	err = service.setClientStatus("123", &ClientStatus{Ok: true})
	require.NoError(t, err)

	err = service.setClientStatus("foo", &ClientStatus{Ok: true})
	require.NoError(t, err)

	err = service.setClientStatus("bar", &ClientStatus{Ok: false, Message: "unable to apply update"})
	require.NoError(t, err)

	clients, err := service.GetClients()
	require.NoError(t, err)

	assert.Len(t, clients, 3)

	// we don't care about order, redis sets are unsorted
	sort.Slice(clients, func(i, j int) bool {
		return clients[i].Id < clients[j].Id
	})

	expected := []*ClientStatus{
		{"123", true, "", nil},
		{"bar", false, "unable to apply update", nil},
		{"foo", true, "", nil},
	}

	for i, s := range clients {
		assert.True(t, s.equal(expected[i]))
	}
}

func TestService_SetLatestFileVersion(t *testing.T) {
	database, err := db.NewDb(os.Getenv("REDIS_ADDRESS"), 0)
	require.NoError(t, err)
	require.NotNil(t, database)

	uri, err := url.Parse("https://example.com/foo/")
	require.NoError(t, err)

	service := &Service{db: database, distUrl: uri}

	err = service.SetLatestFileVersion(10)
	require.NoError(t, err)

	v, err := service.getLatestFileVersion()
	require.NoError(t, err)
	assert.Equal(t, 10, v)
}
