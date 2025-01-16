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
	rows, err := s.db.Query(getConfigsQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var cfgs []Config
	for rows.Next() {
		var cfg Config
		err := rows.Scan(
			&cfg.ID,
			&cfg.CreatedAt,
			&cfg.Desc,
			&cfg.Path,
			&cfg.Type,
			&cfg.Device.Hostname,
			&cfg.Device.MGMTAddress,
		)
		if err != nil {
			log.Println(err)
		}
		cfgs = append(cfgs, cfg)
	}
	return cfgs
}

func (s *Store) AddConfig(cfg Config) (int, error) {
	cfgResult, err := s.db.Exec(insertConfigQuery,
		time.Now(),
		cfg.Path,
		cfg.Type,
		cfg.Desc,
		cfg.Device.ID,
	)
	if err != nil {
		return 0, err
	}
	cfgID, err := cfgResult.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("couldn't get last inserted ID %w", err)
	}
	return int(cfgID), nil
}

func (s *Store) AddDevice(dev Device) (int, error) {
	devResult, err := s.db.Exec(insertDevice,
		time.Now(),
		dev.Hostname,
		dev.MGMTAddress,
		dev.Loopback,
		dev.ISO,
	)
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
			intf.Description,
		)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (s *Store) GetDevices() ([]Device, error) {
	rows, err := s.db.Query(getDevicesQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var devs []Device
	for rows.Next() {
		var dev Device
		err := rows.Scan(
			&dev.ID,
			&dev.Hostname,
			&dev.MGMTAddress,
			&dev.Loopback,
			&dev.ISO,
		)
		devs = append(devs, dev)
		if err != nil {
			log.Println(err)
		}
	}
	return devs, nil
}

func (s *Store) GetDeviceByHostname(hostname string) (Device, error) {
	var dev Device
	err := s.db.QueryRow(getDeviceByHostnameQuery, hostname).Scan(
		&dev.ID,
		&dev.Hostname,
		&dev.MGMTAddress,
		&dev.Loopback,
		&dev.ISO,
	)
	if err != nil {
		log.Println(err)
	}
	return dev, nil
}

func (s *Store) GetDevicesInterfaces(dev Device) (Device, error) {
	rows, err := s.db.Query(getDevicesInterfacesByDeviceIDQuery, dev.ID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var intf Interface
		err := rows.Scan(
			&intf.ID,
			&intf.IFD,
			&intf.IFL,
			&intf.VLAN,
			&intf.Address,
			&intf.Role,
			&intf.Description,
		)
		dev.Interfaces = append(dev.Interfaces, intf)
		if err != nil {
			log.Println(err)
		}
	}
	return dev, nil
}
