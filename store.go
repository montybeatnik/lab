package lab

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore() Store {
	os.Remove("./lab.db")

	db, err := sql.Open("sqlite3", "./lab.db")
	if err != nil {
		log.Println(err)
	}
	return Store{db: db}
}

func (s *Store) CreateTable(table string) error {
	query := `--create
CREATE TABLE IF NOT EXISTS configs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at DATETIME,
	cfg_path TEXT,
	desc TEXT
);	
	`
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("couldn't create table %w", err)
	}
	return nil
}

func (s *Store) GetConfigs() []Config {
	query := `-- get many 
SELECT
	id,
	created_at,
	desc,
	path
FROM configs;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var cfgs []Config
	for rows.Next() {
		var cfg Config
		err := rows.Scan(&cfg.ID, cfg.CreatedAt, cfg.Desc, cfg.Path)
		if err != nil {
			log.Println(err)
		}
		cfgs = append(cfgs, cfg)
	}
	return cfgs
}
