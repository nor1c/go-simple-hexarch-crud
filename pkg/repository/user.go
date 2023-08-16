package repository

import (
	"database/sql"
	"sha/pkg/domain"

)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []domain.User
	if rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
