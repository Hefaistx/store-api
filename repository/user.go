package repository

import (
	"database/sql"
	conf "enigma-laundry-app/config"
	m "enigma-laundry-app/model"
	"fmt"
)

type UserRepository interface {
	CreateUser(user m.User) (m.User, error)
	// GetUserById(id int) (m.User, error)
	// GetUsers() ([]m.User, error)
	// UpdateUser(user m.User) (m.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

func (db *userRepository) CreateUser(user m.User) (m.User, error) {
	err := db.db.QueryRow(conf.CreateUserQuery, user.Username, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.UserID)
	if err != nil {
		return m.User{}, fmt.Errorf("error creating user")
	}

	return user, nil
}

// func (db *userRepository) GetUserById(id int) (m.User, error) {
// 	var user m.User
// 	err := db.db.QueryRow(conf.GetUserById, id).Scan(&user.Id, &user.Name, &user.Phone, &user.Address, &user.Created_At, &user.Updated_At)
// 	if err != nil {
// 		return m.User{}, err
// 	}
// 	return user, nil
// }

// func (db *userRepository) GetUsers() ([]m.User, error) {
// 	var users []m.User
// 	rows, err := db.db.Query(conf.GetUsers)
// 	if err != nil {
// 		return []m.User{}, err
// 	}
// 	for rows.Next() {
// 		var user m.User
// 		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Address, &user.Created_At, &user.Updated_At)
// 		if err != nil {
// 			return []m.User{}, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// func (db *userRepository) UpdateUser(user m.User) (m.User, error) {

// 	err := db.db.QueryRow(conf.UpdateUser, user.Id, user.Name, user.Phone, user.Address, user.Created_At, user.Updated_At)
// 	if err != nil {
// 		return m.User{}, fmt.Errorf("error updating user")
// 	}

// 	return user, nil
// }

func (db *userRepository) DeleteUser(id int) error {

	_, err := db.db.Exec(conf.DeleteUserQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting user")
	}

	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
