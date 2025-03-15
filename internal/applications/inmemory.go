package applications

import (
	"errors"
	"sync"
)

type InMemoryStore struct {
	mu   sync.RWMutex
	apps map[string]*App
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		apps: make(map[string]*App),
	}
}

func (s *InMemoryStore) SaveApp(app *App) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.apps[app.ID] = app
	return nil
}

func (s *InMemoryStore) GetAppByID(id string) (*App, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	app, exists := s.apps[id]
	if !exists {
		return nil, errors.New("app not found")
	}
	return app, nil
}

func (s *InMemoryStore) AddAppVersion(appID string, version *AppVersion) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	app, exists := s.apps[appID]
	if !exists {
		return errors.New("app not found")
	}
	// Check for duplicate versions
	for _, v := range app.Versions {
		if v.Version == version.Version {
			return nil // Version already exists, do not add
		}
	}
	app.Versions = append(app.Versions, *version)
	return nil
}

func (s *InMemoryStore) AddAppArtifact(appID string, artifact *AppArtifact) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	app, exists := s.apps[appID]
	if !exists {
		return errors.New("app not found")
	}
	app.Artifacts = append(app.Artifacts, *artifact)
	return nil
}
