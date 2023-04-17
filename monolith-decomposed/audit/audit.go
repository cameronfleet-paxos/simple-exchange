package audit

import (
	"github.com/approved-designs/simple-exchange/monolith-decomposed/model"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/store"
)

type Auditor struct {
	store store.AuditStore
}

func (a Auditor) Audit(event *model.AuditEvent) {
	a.store.AddAuditEvent(event)
}
