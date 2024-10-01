package user

import (
	"database/sql"
	"fmt"

	"github.com/athulmekkoth/go_server.git/types"
)

type Store struct {
	db *sql.DB
}

func NewsStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserbyEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM USER WHERE EMAIL = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u := new(types.User)
	// Iterate through the result set
	for rows.Next() {
		u, err = ScanRowIntoUser(rows) // Assuming ScanRowIntoUser takes a row and fills the User object
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

}

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {

	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, nil
	}
	return user, nil

}
