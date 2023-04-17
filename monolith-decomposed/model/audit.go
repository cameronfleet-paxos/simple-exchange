package model

// AuditEvent represents a single audit event.
type AuditEvent struct {
	ID         string      `json:"id"`
	Action     string      `json:"action"`
	EntityType string      `json:"entity_type"`
	EntityID   string      `json:"entity_id"`
	Details    interface{} `json:"details"`
	Timestamp  int64       `json:"timestamp"`
}
