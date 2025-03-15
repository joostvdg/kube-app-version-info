package applications

import (
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/rs/zerolog/log"
	"time"
)

// ProcessArgoCDApplication processes the given ArgoCD Application resource
func ProcessArgoCDApplication(storage Store, app *v1alpha1.Application) {
	log.Printf("Processing ArgoCD Application: %s", app.Name)

	// Check if the App already exists
	existingApp, err := storage.GetAppByID(app.Name)
	if err != nil {
		// App does not exist, create a new one
		newApp := &App{
			ID:          app.Name,
			Name:        app.Name,
			Labels:      app.Labels,
			FirstSeenAt: time.Now(),
			LastSeenAt:  time.Now(),
		}
		err = storage.SaveApp(newApp)
		if err != nil {
			log.Error().Err(err).Msg("Failed to save new app")
			return
		}
		existingApp = newApp
	} else {
		// Update the LastSeenAt timestamp
		existingApp.LastSeenAt = time.Now()
	}

	if app.Spec.Source != nil {
		// Create the AppVersion
		appVersion := &AppVersion{
			Version:      app.Spec.Source.TargetRevision,
			DiscoveredAt: time.Now(),
			Labels:       app.Labels,
		}

		// Create AppArtifacts for each image in the status.summary.images list
		if app.Status.Summary.Images != nil {
			for _, image := range app.Status.Summary.Images {
				appArtifact := AppArtifact{
					Source:       image,
					ArtifactType: "image",
					DiscoveredAt: time.Now(),
				}
				appVersion.Artifacts = append(appVersion.Artifacts, appArtifact)
			}
		}

		// Create AppArtifacts for each image in the status.summary.images list
		if app.Status.Summary.Images != nil {
			for _, image := range app.Status.Summary.Images {
				appArtifact := AppArtifact{
					Source:       image,
					ArtifactType: "image",
					DiscoveredAt: time.Now(),
				}
				appVersion.Artifacts = append(appVersion.Artifacts, appArtifact)
			}
		}

		// Add the AppVersion to the App
		err = storage.AddAppVersion(existingApp.ID, appVersion)
		if err != nil {
			log.Error().Err(err).Msg("Failed to add app version")
			return
		}
	}

	// Determine the source(s) to use
	var sources []v1alpha1.ApplicationSource
	if app.Spec.Sources != nil && len(app.Spec.Sources) > 0 {
		sources = app.Spec.Sources
	}

	// Process each source
	for _, source := range sources {
		// Create the AppVersion
		appVersion := &AppVersion{
			Version:      source.TargetRevision,
			DiscoveredAt: time.Now(),
			Labels:       app.Labels,
		}

		// Create AppArtifacts for each image in the status.summary.images list
		if app.Status.Summary.Images != nil {
			for _, image := range app.Status.Summary.Images {
				appArtifact := AppArtifact{
					Source:       image,
					ArtifactType: "image",
					DiscoveredAt: time.Now(),
				}
				appVersion.Artifacts = append(appVersion.Artifacts, appArtifact)
			}
		}

		// Create AppArtifacts for each image in the status.summary.images list
		if app.Status.Summary.Images != nil {
			for _, image := range app.Status.Summary.Images {
				appArtifact := AppArtifact{
					Source:       image,
					ArtifactType: "image",
					DiscoveredAt: time.Now(),
				}
				appVersion.Artifacts = append(appVersion.Artifacts, appArtifact)
			}
		}

		// Add the AppVersion to the App
		err = storage.AddAppVersion(existingApp.ID, appVersion)
		if err != nil {
			log.Error().Err(err).Msg("Failed to add app version")
			return
		}
	}

	// Save the updated App
	err = storage.SaveApp(existingApp)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save updated app")
	}
}
