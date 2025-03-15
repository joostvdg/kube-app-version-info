package applications

import "time"

type App struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Labels      map[string]string `json:"labels"`
	FirstSeenAt time.Time         `json:"first_seen_at"`
	LastSeenAt  time.Time         `json:"last_seen_at"`
	Versions    []AppVersion      `json:"versions"`
	Artifacts   []AppArtifact     `json:"artifacts"`
}

// AppVersion represents a specific version of an application
type AppVersion struct {
	Version      string            `json:"version"`
	DiscoveredAt time.Time         `json:"discovered_at"`
	Labels       map[string]string `json:"labels"`
	Artifacts    []AppArtifact     `json:"artifacts"`
}

// AppArtifact represents an artifact related to an application
type AppArtifact struct {
	Source       string    `json:"source"`
	ArtifactType string    `json:"artifact_type"`
	DiscoveredAt time.Time `json:"discovered_at"`
}
