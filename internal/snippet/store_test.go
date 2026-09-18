package snippet

import (
	"database/sql"
	"errors"
	"os"
	"testing"
)

// Each test gets its own temporary folder, so tests never share kin.db.
func useTempDatabase(t *testing.T) func() {
	t.Helper()

	oldDirectory, err := os.Getwd() // here getwd give the current directory
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	return func() {
		if err := os.Chdir(oldDirectory); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSnippetStorage(t *testing.T) {
	restoreDirectory := useTempDatabase(t)
	defer restoreDirectory()

	first, err := Add("git status")
	if err != nil {
		t.Fatal(err)
	}
	second, err := Add("go test ./...")
	if err != nil {
		t.Fatal(err)
	}

	if first.ID != 1 || second.ID != 2 {
		t.Fatalf("unexpected IDs: got %d and %d", first.ID, second.ID)
	}

	snippets, err := List()
	if err != nil {
		t.Fatal(err)
	}
	if len(snippets) != 2 || snippets[0].Command != "git status" {
		t.Fatalf("unexpected stored snippets: %#v", snippets)
	}

	if err := Remove(first.ID); err != nil {
		t.Fatal(err)
	}
	snippets, err = List()
	if err != nil {
		t.Fatal(err)
	}
	if len(snippets) != 1 || snippets[0].ID != second.ID {
		t.Fatalf("remove did not leave the expected snippet: %#v", snippets)
	}

	if err := Remove(999); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
}
