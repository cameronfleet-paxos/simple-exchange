package audit_test

import (
	"testing"

	"github.com/approved-designs/simple-exchange/monolith-decomposed/audit"
	"github.com/google/go-cmp/cmp"
)

func TestAudit(t *testing.T) {
	auditor := audit.NewAuditor()
	expected := &audit.AuditEvent{
		ID:         "abc",
		Action:     "NEW",
		EntityType: "ORDER",
		EntityID:   "xyz",
		Details:    "something!",
		Timestamp:  11111,
	}
	auditor.Audit(expected)

	actual, err := auditor.GetAudit("abc")

	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Error(diff)
	}
}
