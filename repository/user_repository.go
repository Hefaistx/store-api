package repository

import (
	"database/sql"
	"fmt"
	conf "tokocikbosapi/config"
	m "tokocikbosapi/model"
)

type UserRepository interface {
	CreateUser(user m.User) (m.User, error)
	FindUserByUsernamePasswordQuery(username, password string) (m.UserCredential, error)
	// GetUserById(id int) (m.User, error)
	// GetUsers() ([]m.User, error)
	// UpdateUser(user m.User) (m.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

func (db *userRepository) CreateUser(user m.User) (m.User, error) {
	err := db.db.QueryRow(conf.CreateUserQuery, user.FullName, user.Email, user.Phone).Scan(&user.ID)
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

func (db *userRepository) FindUserByUsernamePasswordQuery(username, password string) (m.UserCredential, error) {
	var cred m.UserCredential

	err := db.db.QueryRow(conf.FindUserByUsernamePasswordQuery, username, password).Scan(&cred.ID, &cred.UserID, &cred.Username, &cred.Password, &cred.Roles, &cred.CreatedAt, &cred.UpdatedAt)
	if err != nil {
		return m.UserCredential{}, err
	}

	return cred, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
