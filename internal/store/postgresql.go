package store

import "github.com/joostvdg/kube-app-version-info/internal/applications"

// PostgreSQLStore implements Store with PostgreSQL
type PostgreSQLStore struct {
	// DB connection
}

func (p *PostgreSQLStore) SaveAppVersion(app *applications.AppVersion) error {
	// Implement DB save logic
	return nil
}

func (p *PostgreSQLStore) GetAppVersions() ([]applications.AppVersion, error) {
	// Implement DB fetch logic
	return nil, nil
}
