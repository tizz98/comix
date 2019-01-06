package cnc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/tizz98/comix/db"
)

type ClientStatus struct {
	Id         string
	Ok         bool
	Message    string
	LastUpdate *time.Time
}

// this does not compare the LastUpdateField
func (s *ClientStatus) equal(other *ClientStatus) bool {
	return s.Id == other.Id && s.Ok == other.Ok && s.Message == other.Message
}

type Service struct {
	db      *db.Db
	distUrl *url.URL
}

func NewTime(t time.Time) *time.Time {
	return &t
}

func New() (*Service, error) {
	addr := viper.GetString("RedisAddress")
	dbNum := viper.GetInt("RedisDbNumber")

	distUrl, err := url.Parse(viper.GetString("UpdateDistributionUrl"))
	if err != nil {
		return nil, err
	}

	database, err := db.NewDb(addr, dbNum)
	if err != nil {
		return nil, err
	}

	return &Service{db: database, distUrl: distUrl}, nil
}

func (s *Service) UpdateStatus(ctx context.Context, update *UpdateMsg) (*Status, error) {
	if !update.UpdateComplete {
		return nil, s.setClientStatus(update.GetClientId(), &ClientStatus{Ok: false, Message: update.GetUpdateMessage()})
	}

	return &Status{}, s.setClientStatus(update.GetClientId(), &ClientStatus{Ok: true})
}

func (s *Service) Ping(ctx context.Context, msg *PingMsg) (*Response, error) {
	if !msg.Ok {
		return nil, s.setClientStatus(msg.GetClientId(), &ClientStatus{Ok: false, Message: msg.GetStatusMessage()})
	}

	if err := s.setClientStatus(msg.GetClientId(), &ClientStatus{Ok: true}); err != nil {
		return nil, err
	}

	if needsNewVersion, err := s.clientNeedsNewVersion(msg.GetClientId()); err != nil {
		return nil, err
	} else if needsNewVersion {
		latestVersion, err := s.getLatestFileVersion()
		if err != nil {
			return nil, err
		}

		sha, err := s.binaryChecksum(latestVersion)
		if err != nil {
			return nil, err
		}

		return &Response{HasUpdate: true, Url: s.binaryUrl(latestVersion), Checksum: sha}, nil
	}

	return &Response{HasUpdate: false}, nil
}

func (s *Service) SetLatestFileVersion(version int) error {
	if _, err := s.db.Set("latest_version", version); err != nil {
		return err
	}
	return nil
}

func (s *Service) getLatestFileVersion() (int, error) {
	var v int

	if _, err := s.db.Get("latest_version", &v); err != nil {
		return 0, err
	}
	return v, nil
}

func (s *Service) setClientStatus(clientId string, status *ClientStatus) error {
	if err := s.addClientId(clientId); err != nil {
		return err
	}

	status.Id = clientId
	status.LastUpdate = NewTime(time.Now())

	if _, err := s.db.Set(s.clientKey(clientId, "status"), status); err != nil {
		return err
	}
	return nil
}

func (s *Service) getClientStatus(clientId string) (*ClientStatus, error) {
	var status *ClientStatus

	if _, err := s.db.Get(s.clientKey(clientId, "status"), &status); err != nil {
		return nil, err
	}
	return status, nil
}

func (s *Service) getClientCurrentVersion(clientId string) (int, error) {
	var v int

	if _, err := s.db.Get(s.clientKey(clientId, "version"), &v); err != nil {
		return 0, err
	}
	return v, nil
}

func (s *Service) setClientCurrentVersion(clientId string, version int) error {
	if err := s.addClientId(clientId); err != nil {
		return err
	}
	if _, err := s.db.Set(s.clientKey(clientId, "version"), version); err != nil {
		return err
	}
	return nil
}

func (s *Service) clientKey(id, key string) string {
	return fmt.Sprintf("%s:%s", id, key)
}

func (s *Service) binaryUrl(version int) string {
	u := *s.distUrl
	u.Path = strings.TrimRight(u.Path, "/")
	u.Path += fmt.Sprintf("/comix-pi/%d", version)
	return u.String()
}

func (s *Service) binaryChecksum(version int) (string, error) {
	u := *s.distUrl
	u.Path = strings.TrimRight(u.Path, "/")
	u.Path += fmt.Sprintf("/comix-pi/%d.sha", version)
	return getChecksum(u.String())
}

func (s *Service) clientNeedsNewVersion(clientId string) (bool, error) {
	currentVersion, err := s.getClientCurrentVersion(clientId)
	if err != nil {
		return false, err
	}

	latestVersion, err := s.getLatestFileVersion()
	if err != nil {
		return false, err
	}

	return latestVersion > currentVersion, nil
}

func (s *Service) getClientIds() ([]string, error) {
	return s.db.SMembers("client_ids")
}

func (s *Service) addClientId(id string) error {
	return s.db.SAdd("client_ids", id)
}

var getChecksum = func(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if b, err := ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	} else {
		return string(b), err
	}
}

func (s *Service) GetClients() ([]*ClientStatus, error) {
	ids, err := s.getClientIds()
	if err != nil {
		return nil, err
	}

	statuses := make([]*ClientStatus, len(ids))

	for i, id := range ids {
		if status, err := s.getClientStatus(id); err != nil {
			return nil, err
		} else {
			statuses[i] = status
		}
	}

	return statuses, err
}
