package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DataBase struct {
	db *sql.DB
}

func ConnectDB() (*DataBase, error) {
	//Conn for MySQL --> conn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	conn := "postgresql://postgres:venky@localhost:5432/project?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Printf("error connecting database %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error ping database %v", err)
		return nil, err
	}
	return &DataBase{db: db}, nil
}

func (d *DataBase) Close() {
	d.db.Close()
}

func (d *DataBase) GetDB() *sql.DB {
	return d.db
}
