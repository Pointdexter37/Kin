package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/Pointdexter37/kin/internal/snippet"
)

// The CLI uses a relative database path, so run it inside a fresh folder.
func useTempCLI(t *testing.T) func() {
	t.Helper()
	oldDirectory, err := os.Getwd()
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

func TestSnippetCommands(t *testing.T) {
	restoreDirectory := useTempCLI(t)
	defer restoreDirectory()

	// SetArgs simulates what a user types after the executable name.
	rootCmd.SetArgs([]string{"add", "git status"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	rootCmd.SetArgs([]string{"list"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	snippets, err := snippet.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(snippets) != 1 || snippets[0].Command != "git status" {
		t.Fatalf("add/list commands failed: %#v", snippets)
	}

	rootCmd.SetArgs([]string{"remove", "1"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
	snippets, err = snippet.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(snippets) != 0 {
		t.Fatalf("remove command failed: %#v", snippets)
	}
}

func TestAddRequiresOneArgument(t *testing.T) {
	rootCmd.SetArgs([]string{"add"})
	if err := rootCmd.Execute(); err == nil {
		t.Fatal("expected add without a command to fail")
	}
}

func TestRemoveValidatesID(t *testing.T) {
	restoreDirectory := useTempCLI(t)
	defer restoreDirectory()

	expectedMessages := map[string]string{
		"abc": "snippet ID must be a number",
		"0":   "snippet ID must be greater than zero",
		"999": "snippet 999 was not found",
	}
	for id, expected := range expectedMessages {
		rootCmd.SetArgs([]string{"remove", id})
		err := rootCmd.Execute()
		if err == nil || !strings.Contains(err.Error(), expected) {
			t.Fatalf("expected remove %q to fail", id)
		}
	}
}
