package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func IsDuplicateEmail(row *sql.Row) (string, error) {
	var email string
	err := row.Scan(&email)
	if err == sql.ErrNoRows {
		return "", nil // not a duplicate
	} else if err != nil {
		return "", err // some error occurred
	}
	return email, nil // duplicate email found
}
