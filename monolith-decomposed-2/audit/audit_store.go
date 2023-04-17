package audit

import "fmt"

type auditStore interface {
	addAuditEvent(event *AuditEvent) error
	getAudit(id string) (*AuditEvent, error)
}

type inMemoryAuditStore struct {
	audits map[string]*AuditEvent
}

func newInMemoryAuditStore() *inMemoryAuditStore {
	return &inMemoryAuditStore{
		audits: make(map[string]*AuditEvent),
	}
}

func (db *inMemoryAuditStore) addAuditEvent(event *AuditEvent) error {
	if _, ok := db.audits[event.ID]; ok {
		return fmt.Errorf("audit with ID '%s' already exists", event.ID)
	}
	db.audits[event.ID] = event
	return nil
}

func (db *inMemoryAuditStore) getAudit(id string) (*AuditEvent, error) {
	event, ok := db.audits[id]
	if !ok {
		return nil, fmt.Errorf("audit with ID '%s' does not exist", id)
	}
	return event, nil
}
