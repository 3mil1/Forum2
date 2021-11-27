package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

type Storage struct {
	config   *Config
	db       *sql.DB
	UserRepo *UserRepo
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("sqlite3", storage.config.DatabaseURI)

	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Connection to db successfully")
	return nil
}

func (storage *Storage) Close() {
	storage.db.Close()
}

func (storage *Storage) AddTables() {
	createSQL, err := ioutil.ReadFile("./createTables.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = storage.db.Exec(string(createSQL))
	if err != nil {
		log.Fatal(err)
	}
}

// User Public repo for user
func (storage *Storage) User() *UserRepo {
	if storage.UserRepo != nil {
		return storage.UserRepo
	}

	storage.UserRepo = &UserRepo{
		storage: storage,
	}

	return storage.UserRepo
}
