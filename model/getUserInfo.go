package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	EMAIL    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func GetUserInfo(username string) (*UserInfo, error) {
	db, err := sql.Open("mysql", "TakumiNakagawara:@Seraio12@(127.0.0.1:3306)/user_information?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user UserInfo
	err = db.QueryRow("SELECT id, name, email, username, password FROM user_information WHERE username = ?", username).Scan(&user.ID, &user.Name, &user.EMAIL, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Will delete soon. It's just for test
func GetAllUsers() ([]UserInfo, error) {
	db, err := sql.Open("mysql", "TakumiNakagawara:@Seraio12@(127.0.0.1:3306)/user_information?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []UserInfo
	rows, err := db.Query("SELECT id, name, username, password FROM user_information")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserInfo
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
