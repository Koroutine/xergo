package xero

import (
	"strings"
	"testing"
)

func TestGetIdempotentKey(t *testing.T) {
	tenantId := "tenant"
	entity := "invoices"
	idempotencyKey := getIdempotencyKey(tenantId, entity)

	if idempotencyKey == "" {
		t.Error("expected idempotencyKey to not be empty")
	}

	if strings.HasPrefix(idempotencyKey, "tenant:invoices:") == false {
		t.Errorf("expected idempotencyKey to start with %s", tenantId)
	}
}
