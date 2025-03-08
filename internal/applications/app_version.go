package applications

import "time"

// AppVersion represents an application and its version tracking
type AppVersion struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Namespace    string    `json:"namespace"`
	Version      string    `json:"version"`
	Source       string    `json:"source"` // e.g., "ArgoCD", "Helm"
	DiscoveredAt time.Time `json:"discovered_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FirstSeenAt  time.Time `json:"first_seen_at"`
	LastSeenAt   time.Time `json:"last_seen_at"`
}
