package belajar_golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,name,email,balance,rating,birth_date,married) values('isro','isro','isro@gmail.com', 1000000, 5.0, NULL, true)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success create new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "select id, name, email, balance, rating, birth_date, married, created_at from customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate, createdAt sql.NullTime
		var married sql.NullBool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		if balance.Valid {
			fmt.Println("Balance : ", balance.Int32)
		}
		if rating.Valid {
			fmt.Println("Rating : ", rating.Float64)
		}
		if birthDate.Valid {
			fmt.Println("Birth Date : ", birthDate.Time)
		}
		if married.Valid {
			fmt.Println("Married : ", married.Bool)
		}
		if createdAt.Valid {
			fmt.Println("Created At : ", createdAt.Time)
		}
		fmt.Println("=================")
	}

}

func TestQueryParam(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	sqlQuery := "SELECT * FROM user where username = ? AND password = ? LIMIT 1"
	// sqlQuery := "SELECT * FROM user where username = '" + username +
	// 	 "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username, password string

		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login : ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	email := "roozy2@gmail.com"
	comment := "Test Comment"

	script := "insert into comments(email,comment) values (?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("success insert new comment with id", insertId)
}

func TestPrepareStatementExec(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSQL := "INSERT INTO comments (email,comment) VALUES(?,?)"

	stmt := CreatePrepareStatement(ctx, db, scriptSQL)

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "isro" + strconv.Itoa(i) + "@gmail.com"
		comment := "Test Comment-" + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
}

type Comment struct {
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

func TestPrepareStatementQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSQL := "SELECT email,comment FROM comments"

	stmt := CreatePrepareStatement(ctx, db,
		scriptSQL)

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		rows.Scan(&c.Email, &c.Comment)
		comments = append(comments, c)
	}
	fmt.Println(comments)
}

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// stmt := CreatePrepareStatement(ctx, db, "INSERT INTO comments (email,comment) VALUES(?,?)")

	// defer stmt.Close()

	script := "INSERT INTO comments (email,comment) VALUES(?,?)"

	for i := 0; i < 10; i++ {
		email := "isro" + strconv.Itoa(i) + "@gmail.com"
		comment := "Test Comment-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}