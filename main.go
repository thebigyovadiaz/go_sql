package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type userStruct struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

type detailUser interface {
	logDetailUser()
}

func (u *userStruct) logDetailUser() {
	fmt.Printf(
		"userId: %d\nusername: %s\npasswprd: %s\ncreatedAt: %s\n\n",
		u.id, u.username, u.password, u.createdAt)
}

func main() {
	db, err := sql.Open("mysql", "root:admin@(127.0.0.1:2223)/test?parseTime=true")
	if err != nil {
		fmt.Printf("Error to connect with database: %v\n", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error to PING DB: %v\n", err.Error())
	}

	fmt.Println("Connection to DB is OK!")

	// Create Table
	/* query := `
		CREATE TABLE users (
		    id INT AUTO_INCREMENT,
		    username TEXT NOT NULL,
		    password TEXT NOT NULL,
		    created_at DATETIME,
		    PRIMARY KEY (id)
	);`

	response, errExec := db.Exec(query)
	if errExec != nil {
		fmt.Printf("Created table error: %v\n", errExec.Error())
	}

	fmt.Println("Users table created in DB") */

	// Insert a user in Database
	/*username := "metaseltrozo"
	password := "secret2024"
	createAt := time.Now()

	query := `INSERT INTO users (username, password, created_at) VALUES (?,?,?)`
	result, errInsert := db.Exec(query, username, password, createAt)
	if errInsert != nil {
		fmt.Printf("Insert user error: %v\n", errInsert.Error())
	}

	userId, errResult := result.LastInsertId()
	if errResult != nil {
		fmt.Printf("Error to get UserID: %v\n", errResult.Error())
	}

	fmt.Printf("User inserted in DB: %v\n", userId)*/

	// Select a user in DB
	query := `SELECT * FROM users WHERE id = ?`
	userId := 2

	var u userStruct
	errSelect := db.QueryRow(query, userId).Scan(
		&u.id,
		&u.username,
		&u.password,
		&u.createdAt)

	if errSelect != nil {
		fmt.Printf("Error to get user information: %v\n", errSelect.Error())
	}

	u.logDetailUser()

	// Select all users
	rows, errAll := db.Query(`SELECT * FROM users`)
	defer rows.Close()

	if errAll != nil {
		fmt.Printf("Error to get all users information: %v\n", errAll.Error())
	}

	var users []userStruct
	for rows.Next() {
		var user userStruct

		errR := rows.Scan(&user.id, &user.username, &user.password, &user.createdAt)
		if errR != nil {
			fmt.Printf("Error to insert information in array: %v\n", errR.Error())
		}

		users = append(users, user)
	}

	errRow := rows.Err()
	if errRow != nil {
		fmt.Printf("Error to get DB information: %v\n", errRow.Error())
	}

	fmt.Println("List All Users:")

	for _, user := range users {
		user.logDetailUser()
	}

	fmt.Println("Information gotten successfully!")

	// Delete a user
	queryDelete := `DELETE FROM users WHERE id = ?`

	_, errDelete := db.Exec(queryDelete, 1)
	if errDelete != nil {
		fmt.Printf("Error to delete the userID: %d - Error: %v\n", 1, errRow.Error())
	}

	fmt.Println("User deleted successfully!")
}
