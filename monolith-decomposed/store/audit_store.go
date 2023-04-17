package store

import "github.com/approved-designs/simple-exchange/monolith-decomposed/model"

type AuditStore interface {
	AddAuditEvent(event *model.AuditEvent)
}

type InMemoryAuditStore struct {
	audits []*model.AuditEvent
}

func (db *InMemoryAuditStore) AddAuditEvent(event *model.AuditEvent) {
	db.audits = append(db.audits, event)
}
