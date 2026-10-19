package http

import (
	"strings"
	"testing"
)

func TestETagForDeterministic(t *testing.T) {
	e1 := ETagFor("users", `{"id":1}`)
	e2 := ETagFor("users", `{"id":1}`)
	if e1 != e2 {
		t.Fatalf("same content should yield same etag: %s vs %s", e1, e2)
	}
	if e1[0:3] != `W/"` {
		t.Fatalf("etag should be weak form: %s", e1)
	}
	if ETagFor("orders", `{"id":1}`) == e1 {
		t.Fatal("different namespaces should yield different etags")
	}
}

func TestETagMatches(t *testing.T) {
	etag := ETagFor("ns", "abc")
	weak := etag
	strong := strings.TrimPrefix(etag, "W/")

	tests := []struct {
		name      string
		imm       string
		wantMatch bool
	}{
		{"empty", "", false},
		{"star", "*", true},
		{"exact", weak, true},
		{"strong matches weak", strong, true},
		{"list hit", `W/"x", ` + strong, true},
		{"list miss", `W/"x", W/"y"`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ETagMatches(tt.imm, etag); got != tt.wantMatch {
				t.Fatalf("ETagMatches(%q, etag) = %v, want %v", tt.imm, got, tt.wantMatch)
			}
		})
	}
}
