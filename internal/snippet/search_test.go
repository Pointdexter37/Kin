package snippet

import "testing"

func TestSearchFindsTextWithoutCaseSensitivity(t *testing.T) {
	restoreDirectory := useTempDatabase(t)
	defer restoreDirectory()

	if _, err := Add("git status"); err != nil {
		t.Fatal(err)
	}
	if _, err := Add("docker ps"); err != nil {
		t.Fatal(err)
	}

	results, err := Search("GIT")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0].Command != "git status" {
		t.Fatalf("unexpected search results: %#v", results)
	}
}
