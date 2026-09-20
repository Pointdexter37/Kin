package snippet

import (
	"database/sql"
	"strings"

	// Importing the driver registers SQLite with database/sql.
	// The blank identifier means this package is used for its setup code.
	_ "modernc.org/sqlite"
)

const databaseFile = "kin.db"

// openDatabase opens the SQLite file and makes sure the table exists.
// SQLite creates the file automatically when it does not exist yet.
func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", databaseFile)
	if err != nil {
		return nil, err
	}

	// The CLI opens the database for one short operation at a time.
	// Closing idle connections helps Windows release the database file quickly.
	db.SetMaxIdleConns(0)

	// CREATE TABLE IF NOT EXISTS is safe to run every time the CLI starts.
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS snippets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			command TEXT NOT NULL,
			tags TEXT NOT NULL DEFAULT ''
		)
	`)
	if err != nil {
		db.Close()
		return nil, err
	}

	// Older databases created before tag support will not have the `tags` column.
	// SQLite allows us to add the column safely when it is missing.
	_, err = db.Exec(`ALTER TABLE snippets ADD COLUMN tags TEXT NOT NULL DEFAULT ''`)
	if err != nil && !isColumnExistsError(err) {
		db.Close()
		return nil, err
	}

	return db, nil
}

func isColumnExistsError(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "duplicate column name") || strings.Contains(err.Error(), "already exists"))
}

// Init creates the database and table without adding a snippet.
func Init() error {
	db, err := openDatabase()
	if err != nil {
		return err
	}
	return db.Close()
}

// Add stores one command and returns the complete saved snippet.
// The tags are stored as a single comma-separated string to keep persistence simple.
func Add(command string, tags ...string) (Snippet, error) {
	db, err := openDatabase()
	if err != nil {
		return Snippet{}, err
	}
	defer db.Close()

	joinedTags := strings.Join(tags, ",")
	result, err := db.Exec("INSERT INTO snippets (command, tags) VALUES (?, ?)", command, joinedTags)
	if err != nil {
		return Snippet{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Snippet{}, err
	}
	return Snippet{ID: int(id), Command: command, Tags: joinedTags}, nil
}

// List reads snippets in the order they were added.
func List() ([]Snippet, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, command, tags FROM snippets ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var item Snippet
		if err := rows.Scan(&item.ID, &item.Command, &item.Tags); err != nil {
			return nil, err
		}
		snippets = append(snippets, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}

// Get returns one snippet by its saved ID.
func Get(id int) (Snippet, error) {
	db, err := openDatabase()
	if err != nil {
		return Snippet{}, err
	}
	defer db.Close()

	var item Snippet
	err = db.QueryRow(
		"SELECT id, command, tags FROM snippets WHERE id = ?", id,
	).Scan(&item.ID, &item.Command, &item.Tags)
	return item, err
}

// Search returns snippets whose commands or tags contain the query text.
// SQLite's NOCASE makes normal English letters match upper/lower case.
func Search(query string) ([]Snippet, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	searchValue := "%" + query + "%"
	rows, err := db.Query(
		"SELECT id, command, tags FROM snippets WHERE command LIKE ? COLLATE NOCASE OR tags LIKE ? COLLATE NOCASE ORDER BY id",
		searchValue,
		searchValue,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var item Snippet
		if err := rows.Scan(&item.ID, &item.Command, &item.Tags); err != nil {
			return nil, err
		}
		snippets = append(snippets, item)
	}
	return snippets, rows.Err()
}

// Remove deletes exactly one snippet by its database ID.
func Remove(id int) error {
	db, err := openDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM snippets WHERE id = ?", id)
	if err != nil {
		return err
	}

	removed, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if removed == 0 {
		return sql.ErrNoRows
	}
	return nil
}
