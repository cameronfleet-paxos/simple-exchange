package audit

type AuditEvent struct {
	ID         string      `json:"id"`
	Action     string      `json:"action"`
	EntityType string      `json:"entity_type"`
	EntityID   string      `json:"entity_id"`
	Details    interface{} `json:"details"`
	Timestamp  int64       `json:"timestamp"`
}

type Auditor struct {
	store auditStore
}

func NewAuditor() *Auditor {
	return &Auditor{
		store: newInMemoryAuditStore(),
	}
}

func (a *Auditor) Audit(event *AuditEvent) error {
	return a.store.addAuditEvent(event)
}

func (a *Auditor) GetAudit(id string) (*AuditEvent, error) {
	return a.store.getAudit(id)
}
