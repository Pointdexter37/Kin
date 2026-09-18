package cmd

import "testing"

func TestSearchRequiresText(t *testing.T) {
	rootCmd.SetArgs([]string{"search"})
	if err := rootCmd.Execute(); err == nil {
		t.Fatal("expected search without text to fail")
	}
}
