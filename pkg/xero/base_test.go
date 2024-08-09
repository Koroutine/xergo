package xero

import (
	"strings"
	"testing"
)

func TestGetIdempotentKey(t *testing.T) {
	tenantId := "tenant"
	entity := "invoices"
	idempotencyKey, err := getIdempotencyKey(tenantId, entity)

	if err != nil {
		t.Error("expected no error")
	}

	if idempotencyKey == "" {
		t.Error("expected idempotencyKey to not be empty")
	}

	if strings.HasPrefix(idempotencyKey, "tenant:invoices:") == false {
		t.Errorf("expected idempotencyKey to start with %s", tenantId)
	}
}
