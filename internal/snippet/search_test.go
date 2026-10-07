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

func TestSearchMatchesAllWordsAndRanksCommandMatchesFirst(t *testing.T) {
	restoreDirectory := useTempDatabase(t)
	defer restoreDirectory()

	first, err := Add("release production", "deploy")
	if err != nil {
		t.Fatal(err)
	}
	second, err := Add("git status", "common")
	if err != nil {
		t.Fatal(err)
	}
	third, err := Add("show notes", "release")
	if err != nil {
		t.Fatal(err)
	}

	results, err := Search("GIT STATUS")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0].ID != second.ID {
		t.Fatalf("expected only the multi-word command match, got %#v", results)
	}

	results, err = Search("release")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 || results[0].ID != first.ID || results[1].ID != third.ID {
		t.Fatalf("expected command match before tag match, got %#v", results)
	}
}

func TestSearchIgnoresWhitespaceOnlyQueries(t *testing.T) {
	restoreDirectory := useTempDatabase(t)
	defer restoreDirectory()

	if _, err := Add("git status"); err != nil {
		t.Fatal(err)
	}

	results, err := Search("   ")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 0 {
		t.Fatalf("expected no results for whitespace-only query, got %#v", results)
	}
}
