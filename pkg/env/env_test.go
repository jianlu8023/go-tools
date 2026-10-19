package env

import (
	"testing"
	"time"
)

func TestGetInt(t *testing.T) {
	t.Setenv("GO_TOOLS_TEST_INT", "")
	if got := GetInt("GO_TOOLS_TEST_INT", 7); got != 7 {
		t.Fatalf("empty env should return default, got %d", got)
	}
	t.Setenv("GO_TOOLS_TEST_INT", "42")
	if got := GetInt("GO_TOOLS_TEST_INT", 7); got != 42 {
		t.Fatalf("expected 42, got %d", got)
	}
	t.Setenv("GO_TOOLS_TEST_INT", "not-a-number")
	if got := GetInt("GO_TOOLS_TEST_INT", 7); got != 7 {
		t.Fatalf("invalid env should return default, got %d", got)
	}
}

func TestGetString(t *testing.T) {
	t.Setenv("GO_TOOLS_TEST_STR", "hello")
	if got := GetString("GO_TOOLS_TEST_STR", "d"); got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
	if got := GetString("GO_TOOLS_TEST_STR_MISSING", "d"); got != "d" {
		t.Fatalf("expected default, got %s", got)
	}
}

func TestGetBool(t *testing.T) {
	t.Setenv("GO_TOOLS_TEST_BOOL", "1")
	if got := GetBool("GO_TOOLS_TEST_BOOL", false); !got {
		t.Fatal("expected true for '1'")
	}
	t.Setenv("GO_TOOLS_TEST_BOOL", "true")
	if got := GetBool("GO_TOOLS_TEST_BOOL", false); !got {
		t.Fatal("expected true for 'true'")
	}
	t.Setenv("GO_TOOLS_TEST_BOOL", "x")
	if got := GetBool("GO_TOOLS_TEST_BOOL", true); !got {
		t.Fatal("invalid value should return default true")
	}
}

func TestGetDuration(t *testing.T) {
	t.Setenv("GO_TOOLS_TEST_DUR", "1m30s")
	if got := GetDuration("GO_TOOLS_TEST_DUR", time.Second); got != 90*time.Second {
		t.Fatalf("expected 1m30s, got %v", got)
	}
	if got := GetDuration("GO_TOOLS_TEST_DUR_MISSING", 5*time.Second); got != 5*time.Second {
		t.Fatalf("expected default, got %v", got)
	}
}
