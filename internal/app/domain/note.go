package domain

type Note struct {
	ID          uint   `db:"id"`
	Name        string `db:"name"`
	Link        string `db:"link"`
	Description string `db:"description"`
}
