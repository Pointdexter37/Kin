package snippet

var snippets []Snippet // temp storage for snippets

func Add(command string) Snippet {
	// if we do snippet.Add("git status") we create snippet ID : 1 and command: "git status
	snippet := Snippet {
		ID : len(snippets) + 1,
		Command: command,
	}
	snippets = append(snippets, snippet)
	//puts it inot our slice

	return snippet
}

func List() []Snippet {
	return snippets
	// We're allowing the caller to read our stored 
	// snippets without directly accessing the snippets variable.
}



