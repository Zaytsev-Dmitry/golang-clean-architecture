package queries

const (
	InsertNoteQuery = "INSERT INTO notes (name, link, description) VALUES ($1, $2, $3) RETURNING *"
	SelectNoteByID  = "SELECT * FROM notes WHERE id = $1"
)
