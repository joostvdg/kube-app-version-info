package store

import "github.com/joostvdg/kube-app-version-info/internal/applications"

// Store interface abstracts database interactions
type Store interface {
	SaveAppVersion(app *applications.AppVersion) error
	GetAppVersions() ([]applications.AppVersion, error)
}
