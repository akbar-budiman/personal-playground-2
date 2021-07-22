package service

import (
	"context"
	"fmt"
	"log"

	"github.com/akbar-budiman/personal-playground-2/entity"
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

	createTableQuery := "CREATE TABLE IF NOT EXISTS users (name VARCHAR PRIMARY KEY, age INT, address VARCHAR, searchable VARCHAR);"
	if _, err := dbpool.Exec(context.Background(), createTableQuery); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Local CRDB DDL initialized")
}

func GetUser(name string) []entity.User {
	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	query := fmt.Sprintf("SELECT * FROM users WHERE name='%s';", name)
	rows, err := dbpool.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.Name, &u.Age, &u.Address, &u.Searchable)
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

func InsertUser(user *entity.User) {
	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	query := fmt.Sprintf(
		"INSERT INTO users(name, age, address, searchable) VALUES('%s', %d, '%s', '%s');",
		user.Name,
		user.Age,
		user.Address,
		user.Searchable,
	)

	if _, err := dbpool.Exec(context.Background(), query); err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(name string, user *entity.User) {
	dbpool, err := pgxpool.ConnectConfig(context.Background(), CrdbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	query := fmt.Sprintf(
		"UPDATE users SET age=%d, address='%s', searchable='%s' WHERE name='%s';",
		user.Age,
		user.Address,
		user.Searchable,
		user.Name,
	)

	if _, err := dbpool.Exec(context.Background(), query); err != nil {
		log.Fatal(err)
	}
}
