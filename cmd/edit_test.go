package cmd

import (
	"testing"

	"github.com/Pointdexter37/kin/internal/snippet"
)

func TestEditUpdatesSnippet(t *testing.T) {
	restoreDirectory := useTempCLI(t)
	defer restoreDirectory()

	if _, err := snippet.Add("echo hello"); err != nil {
		t.Fatal(err)
	}

	rootCmd.SetArgs([]string{"edit", "1", "echo updated", "--tag", "dev", "--tag", "shell"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	item, err := snippet.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if item.Command != "echo updated" || item.Tags != "dev,shell" {
		t.Fatalf("unexpected edited snippet: %#v", item)
	}
}
