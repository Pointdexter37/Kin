package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Pointdexter37/kin/internal/snippet"
)

func TestRunOnlyPreviewsSnippet(t *testing.T) {
	restoreDirectory := useTempCLI(t)
	defer restoreDirectory()

	item, err := snippet.Add("echo safe preview")
	if err != nil {
		t.Fatal(err)
	}

	var output bytes.Buffer
	rootCmd.SetOut(&output)
	t.Cleanup(func() { rootCmd.SetOut(nil) })
	rootCmd.SetArgs([]string{"run", "1"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "Preview (not executed): "+item.Command) {
		t.Fatalf("unexpected preview output: %q", output.String())
	}
}
