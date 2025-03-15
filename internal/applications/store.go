package applications

type Store interface {
	SaveApp(app *App) error
	GetAppByID(id string) (*App, error)
	AddAppVersion(appID string, version *AppVersion) error
	AddAppArtifact(appID string, artifact *AppArtifact) error
}
