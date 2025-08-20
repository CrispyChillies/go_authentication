package models

import (
	"database/sql"
	"fmt"

	"github.com/CrispyChillies/go_authentication/utils"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func RegisterUser(db *sql.DB, username, password string) error {
	hasedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err = db.Exec(query, username, string(hasedPassword))
	if err != nil {
		return err
	}

	fmt.Println("User registered successfully")
	return nil
}
