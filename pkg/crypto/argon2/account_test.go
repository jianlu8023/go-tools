package argon2

import (
	"strings"
	"testing"
)

func TestAccountPasswordRoundTrip(t *testing.T) {
	hash, err := HashAccountPassword("correct-horse")
	if err != nil {
		t.Fatalf("Hash: %v", err)
	}
	if !strings.HasPrefix(hash, "$argon2id$v=19$m=19456,t=2,p=1$") {
		t.Fatalf("unexpected hash format: %s", hash)
	}
	if !VerifyAccountPassword("correct-horse", hash) {
		t.Fatal("verify should accept correct password")
	}
	if VerifyAccountPassword("wrong-password", hash) {
		t.Fatal("verify should reject wrong password")
	}
}

func TestAccountPasswordValidate(t *testing.T) {
	if err := ValidateAccountPassword("short"); err == nil {
		t.Fatal("should reject <8 chars")
	}
	long := "a" + strings.Repeat("b", 129)
	if err := ValidateAccountPassword(long); err == nil {
		t.Fatal("should reject >128 chars")
	}
	if err := ValidateAccountPassword(strings.Repeat("中文密码啊", 4)); err != nil {
		t.Fatalf("valid 12-char utf8 password rejected: %v", err)
	}
}

func TestAccountPasswordVerifyBadFormat(t *testing.T) {
	for _, bad := range []string{
		"",
		"$argon2i$v=19$m=19456,t=2,p=1$x$y",
		"$argon2id$v=19$m=1,t=2,p=1$x$y",
		"garbage",
	} {
		if VerifyAccountPassword("correct-horse", bad) {
			t.Fatalf("should reject malformed hash: %q", bad)
		}
	}
}
