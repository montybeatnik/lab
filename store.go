package lab

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore() Store {
	// os.Remove("./lab.db")

	db, err := sql.Open("sqlite3", "./lab.db")
	if err != nil {
		log.Println(err)
	}
	return Store{db: db}
}

func (s *Store) CreateTable(query string) error {
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

func (s *Store) AddDevice(dev Device) (int, error) {
	devResult, err := s.db.Exec(insertDevice, time.Now(), dev.Hostname, dev.MGMTAddress)
	if err != nil {
		return 0, err
	}
	devID, err := devResult.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("couldn't get last inserted ID %w", err)
	}
	dev.ID = int(devID)
	if err := s.AddInterfaces(dev); err != nil {
		return 0, err
	}
	return dev.ID, nil
}

func (s *Store) AddInterfaces(dev Device) error {
	for _, intf := range dev.Interfaces {
		_, err := s.db.Exec(insertInterface,
			time.Now(),
			dev.ID,
			intf.IFD,
			intf.IFL,
			intf.VLAN,
			intf.Address,
			intf.Role,
		)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (s *Store) GetDevices() ([]Device, error) {
	rows, err := s.db.Query(getDeviceQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var devs []Device
	for rows.Next() {
		var dev Device
		var intf Interface
		err := rows.Scan(
			&dev.ID,
			&dev.Hostname,
			&dev.MGMTAddress,
			&intf.ID,
			&intf.IFD,
			&intf.IFL,
			&intf.VLAN,
			&intf.Address,
			&intf.Role,
		)
		dev.Interfaces = append(dev.Interfaces, intf)
		devs = append(devs, dev)
		if err != nil {
			log.Println(err)
		}
	}
	return devs, nil
}
