package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Pointdexter37/kin/internal/snippet"
)

func TestSearchRequiresText(t *testing.T) {
	rootCmd.SetArgs([]string{"search"})
	if err := rootCmd.Execute(); err == nil {
		t.Fatal("expected search without text to fail")
	}
}

func TestChooseSnippet(t *testing.T) {
	var output bytes.Buffer
	rootCmd.SetOut(&output)
	rootCmd.SetIn(strings.NewReader("2\n"))
	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetIn(nil)
	})

	err := chooseSnippet(rootCmd, []snippet.Snippet{
		{ID: 1, Command: "git status"},
		{ID: 2, Command: "docker ps"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "docker ps") {
		t.Fatalf("expected selected command, got %q", output.String())
	}
}
