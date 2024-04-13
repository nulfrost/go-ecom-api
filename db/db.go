package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteStorage(dbName string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbName)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	_, err = db.Exec("create table if not exists users (id integer not null primary key, first_name text, last_name text, email text, password text, created_at timestamp default current_timestamp not null)")

	if err != nil {
		return nil, err
	}

	return db, nil
}
