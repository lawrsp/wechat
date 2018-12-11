package token

import (
	"fmt"
	"sync"

	"github.com/lawrsp/wechat/config"
)

type AppTokenInfo struct {
	config.AppConfig
	TokenInfo
}

type Server struct {
	lock *sync.RWMutex
	apps map[string]*AppTokenInfo
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RegisterApp(app config.AppConfig) {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.apps[app.AppID] = &AppTokenInfo{AppConfig: app}
}

func (s *Server) UnregisterApp(appID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.apps, appID)
}

func (s *Server) RefreshToken(appID string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()

	app, ok := s.apps[appID]
	if !ok {
		return fmt.Errorf("not registered")
	}

	tokenInfo, err := GetAccessToken(app.AppID, app.AppSecret)
	if err != nil {
		return err
	}

	app.TokenInfo = *tokenInfo
	return nil
}

func (s *Server) StartAutoRefresh() {

}

func (s *Server) StartServer() {

}
