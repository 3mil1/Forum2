package storage

import (
	"fmt"
	"forum/internal/app/models"
	"log"
)

type UserRepo struct {
	storage *Storage
}

var (
	tableUser = "users"
)

// Create User in db
func (ur *UserRepo) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (email, login, password) VALUES ($1, $2, $3) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Email, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//FindByLogin Find user by Login
//func (ur *UserRepo) FindByLogin(login string) (*models.User, bool, error) {
//	return nil, false, nil
//}

func (ur *UserRepo) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Куда читаем
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Email, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
