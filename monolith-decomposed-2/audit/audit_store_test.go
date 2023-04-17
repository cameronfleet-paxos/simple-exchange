package audit

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_addAuditEvent(t *testing.T) {
	store := newInMemoryAuditStore()
	expected := &AuditEvent{
		ID:         "abc",
		Action:     "NEW",
		EntityType: "ORDER",
		EntityID:   "xyz",
		Details:    "something!",
		Timestamp:  11111,
	}
	store.addAuditEvent(expected)

	if len(store.audits) != 1 {
		t.Fatal("Expected only single audit")
	}

	if diff := cmp.Diff(expected, store.audits["abc"]); diff != "" {
		t.Fatal(diff)
	}
}


func Test_getAudit(t *testing.T) {
	// TODO
}