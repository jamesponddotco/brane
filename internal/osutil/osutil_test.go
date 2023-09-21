package osutil_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/brane/internal/osutil"
)

func TestUserDataDir(t *testing.T) { //nolint:paralleltest // t.Setenv is not thread-safe
	path := osutil.UserDataDir()
	if path == "" {
		t.Fatal("expected data dir path, but got empty string")
	}

	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("HOME", "")

	path = osutil.UserDataDir()
	if path != "" {
		t.Fatalf("expected empty string, but got %q", path)
	}
}
