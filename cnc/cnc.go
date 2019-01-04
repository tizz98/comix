package cnc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/tizz98/comix/db"
)

type ClientStatus struct {
	Ok      bool
	Message string
}

type Service struct {
	db      *db.Db
	distUrl *url.URL
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

//func (s *Service) ApplyPatch(ctx context.Context, u *CnCUpdate) (*Status, error) {
//	resp, err := http.Get(u.GetUrl())
//	if err != nil {
//		return &Status{Success: false}, err
//	}
//	defer resp.Body.Close()
//
//	err = update.Apply(resp.Body, update.Options{
//		Hash:     crypto.SHA256,
//		Checksum: u.GetChecksum(),
//	})
//	if err != nil {
//		if rerr := update.RollbackError(err); rerr != nil {
//			logrus.WithError(rerr).Error("failed to rollback from bad update")
//		}
//	}
//
//	return &Status{Success: err == nil}, err
//}

func (s *Service) Ping(ctx context.Context, msg *PingMsg) (*Response, error) {
	if !msg.Ok {
		return nil, s.setClientStatus(msg.GetClientId(), msg.GetStatusMessage())
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

func (s *Service) setLatestFileVersion(version int) error {
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

func (s *Service) setClientStatus(clientId, status string) error {
	if _, err := s.db.Set(s.clientKey(clientId, "status"), status); err != nil {
		return err
	}
	return nil
}

func (s *Service) getClientStatus(clientId string) (string, error) {
	var status string

	if _, err := s.db.Get(s.clientKey(clientId, "status"), &status); err != nil {
		return "", err
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
	u.Path = fmt.Sprintf("/comix-pi/%d", version)
	return u.String()
}

func (s *Service) binaryChecksum(version int) ([]byte, error) {
	u := *s.distUrl
	u.Path = fmt.Sprintf("/comix-pi/%d.sha", version)
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

var getChecksum = func(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
