package fsutil

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFileNew(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")

	action, err := WriteFile(dest, []byte("hello"), Options{})
	if err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}
	if action != ActionOverwrite {
		t.Errorf("action = %v, want ActionOverwrite", action)
	}

	got, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, []byte("hello")) {
		t.Errorf("contents = %q, want hello", got)
	}
}

func TestWriteFileIdenticalContentSkips(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("same"), 0o644); err != nil {
		t.Fatal(err)
	}

	action, err := WriteFile(dest, []byte("same"), Options{})
	if err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}
	if action != ActionSkip {
		t.Errorf("action = %v, want ActionSkip", action)
	}
}

func TestWriteFileNoClobberSkipsOnDiff(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	action, err := WriteFile(dest, []byte("new"), Options{NoClobber: true})
	if err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}
	if action != ActionSkip {
		t.Errorf("action = %v, want ActionSkip", action)
	}
	got, _ := os.ReadFile(dest)
	if !bytes.Equal(got, []byte("old")) {
		t.Errorf("file changed despite NoClobber: %q", got)
	}
}

func TestWriteFileForceOverwrites(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	action, err := WriteFile(dest, []byte("new"), Options{Force: true})
	if err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}
	if action != ActionOverwrite {
		t.Errorf("action = %v, want ActionOverwrite", action)
	}
	got, _ := os.ReadFile(dest)
	if !bytes.Equal(got, []byte("new")) {
		t.Errorf("contents = %q, want new", got)
	}
}

func TestWriteFileOnConflictOverwrite(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	called := false
	opts := Options{
		OnConflict: func(path string, existing, incoming []byte) Action {
			called = true
			if !bytes.Equal(existing, []byte("old")) {
				t.Errorf("existing = %q, want old", existing)
			}
			if !bytes.Equal(incoming, []byte("new")) {
				t.Errorf("incoming = %q, want new", incoming)
			}
			return ActionOverwrite
		},
	}

	action, err := WriteFile(dest, []byte("new"), opts)
	if err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Fatal("OnConflict was not called")
	}
	if action != ActionOverwrite {
		t.Errorf("action = %v, want ActionOverwrite", action)
	}
}

func TestWriteFileOnConflictSkip(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	opts := Options{
		OnConflict: func(path string, existing, incoming []byte) Action {
			return ActionSkip
		},
	}
	action, err := WriteFile(dest, []byte("new"), opts)
	if err != nil {
		t.Fatal(err)
	}
	if action != ActionSkip {
		t.Errorf("action = %v, want ActionSkip", action)
	}
	got, _ := os.ReadFile(dest)
	if !bytes.Equal(got, []byte("old")) {
		t.Errorf("file changed despite skip: %q", got)
	}
}

func TestWriteFileOnConflictAbort(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}

	opts := Options{
		OnConflict: func(path string, existing, incoming []byte) Action {
			return ActionAbort
		},
	}
	action, err := WriteFile(dest, []byte("new"), opts)
	if !errors.Is(err, ErrAborted) {
		t.Fatalf("err = %v, want ErrAborted", err)
	}
	if action != ActionAbort {
		t.Errorf("action = %v, want ActionAbort", action)
	}
}

func TestWriteFileDryRun(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")

	action, err := WriteFile(dest, []byte("hello"), Options{DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if action != ActionOverwrite {
		t.Errorf("action = %v, want ActionOverwrite", action)
	}
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("dry-run wrote file: stat err = %v", err)
	}
}

func TestWriteFileDryRunIdentical(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("same"), 0o644); err != nil {
		t.Fatal(err)
	}
	action, err := WriteFile(dest, []byte("same"), Options{DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if action != ActionSkip {
		t.Errorf("action = %v, want ActionSkip", action)
	}
}

func TestWriteFileCreatesParentDirs(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "nested", "deep", "out.txt")

	if _, err := WriteFile(dest, []byte("hi"), Options{}); err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}
	if _, err := os.Stat(dest); err != nil {
		t.Errorf("dest not written: %v", err)
	}
}

func TestWriteFileTempFileCleanedUpOnSuccess(t *testing.T) {
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if _, err := WriteFile(dest, []byte("hi"), Options{}); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if e.Name() != "out.txt" {
			t.Errorf("unexpected file leftover: %s", e.Name())
		}
	}
}

func TestWriteFileNoClobberAndForceConflict(t *testing.T) {
	// Caller responsibility to set only one - but assert behaviour: NoClobber wins.
	// (Mutually-exclusive validation happens at the cobra layer.)
	dir := t.TempDir()
	dest := filepath.Join(dir, "out.txt")
	if err := os.WriteFile(dest, []byte("old"), 0o644); err != nil {
		t.Fatal(err)
	}
	action, err := WriteFile(dest, []byte("new"), Options{NoClobber: true, Force: true})
	if err != nil {
		t.Fatal(err)
	}
	if action != ActionSkip {
		t.Errorf("action = %v, want ActionSkip (NoClobber wins)", action)
	}
}
