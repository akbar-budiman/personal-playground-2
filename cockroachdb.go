package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	crdbAddress  = "localhost"
	crdbPort     = "26257"
	crdbUsername = "abay"
	crdbPassword = "abaydatabase"
	crdbDatabase = "abaydatabase"
	crdbMaxPool  = "2"
	CrdbConfig   *pgxpool.Config
)

func InitializeLocalCrdbPool() {
	fmt.Println("Initializing local CRDB pool")

	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require&pool_max_conns=%s",
		crdbUsername,
		crdbPassword,
		crdbAddress,
		crdbPort,
		crdbDatabase,
		crdbMaxPool,
	)

	config, err := pgxpool.ParseConfig(conn)
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	CrdbConfig = config

	fmt.Println("Local CRDB pool initialized")
}

func InitializeLocalCrdbDdl() {
	fmt.Println("Initializing local CRDB DDL")

	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	createTableQuery := "CREATE TABLE IF NOT EXISTS users (name VARCHAR PRIMARY KEY, age INT, address VARCHAR);"
	if _, err := dbpool.Exec(context.Background(), createTableQuery); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Local CRDB DDL initialized")
}

func GetUser(name string) []User {
	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	query := fmt.Sprintf("SELECT * FROM users WHERE name='%s';", name)
	fmt.Println("query:", query)
	rows, err := dbpool.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Age, &u.Address)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func InsertUser(user *User) {
	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	query := fmt.Sprintf(
		"INSERT INTO users(name, age, address) VALUES('%s', %d, '%s');",
		user.Name,
		user.Age,
		user.Address,
	)

	if _, err := dbpool.Exec(context.Background(), query); err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(name string, user *User) {
	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	query := fmt.Sprintf(
		"UPDATE users SET age=%d, address='%s' WHERE name='%s';",
		user.Age,
		user.Address,
		user.Name,
	)

	if _, err := dbpool.Exec(context.Background(), query); err != nil {
		log.Fatal(err)
	}
}
