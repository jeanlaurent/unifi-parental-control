package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type kidsDB struct {
	kdb *sql.DB
}

func initDb() (*kidsDB, error) {
	db, err := sql.Open("sqlite3", "./unifi-kids.db")
	if err != nil {
		return nil, err
	}
	kidsDb := &kidsDB{db}
	err = kidsDb.init()
	if err != nil {
		return nil, err
	}
	return kidsDb, nil
}

func (k *kidsDB) init() error {
	statement, err := k.kdb.Prepare("CREATE TABLE IF NOT EXISTS kids (kidId INTEGER PRIMARY KEY, firstname TEXT)")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	statement, err = k.kdb.Prepare("CREATE TABLE IF NOT EXISTS devices (deviceId INTEGER PRIMARY KEY, name TEXT, mac TEXT, kidId INTEGER, FOREIGN KEY(kidId) REFERENCES kids(kidId))")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	statement, err = k.kdb.Prepare("CREATE TABLE IF NOT EXISTS blocked (blockedId INTEGER PRIMARY KEY, name TEXT, mac TEXT)")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	// statement, err = k.kdb.Prepare("CREATE TABLE IF NOT EXISTS schedules (scheduleId INTEGER PRIMARY KEY, start timestamp/datetime, end timestamp/datetime, FOREIGN KEY(kidsId) REFERENCES kids(kidsId))")
	// if err != nil {
	// 	return err
	// }
	// _, err = statement.Exec()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (k *kidsDB) addBlockedDevice(name, mac string) error {
	statement, err := k.kdb.Prepare("INSERT INTO blocked (name, mac) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(name, mac)
	return err
}

func (k *kidsDB) deleteBlockedDevice(id int) error {
	statement, err := k.kdb.Prepare("DELETE FROM blocked WHERE blockedID = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}

func (k *kidsDB) allBlocked() ([]device, error) {
	rows, err := k.kdb.Query("SELECT id, name, mac FROM blocked")
	if err != nil {
		return nil, err
	}
	var id int
	var name string
	var mac string
	var devices []device
	for rows.Next() {
		err := rows.Scan(&id, &name, &mac)
		if err != nil {
			return devices, err
		}
		machine := device{id, name, mac}
		devices = append(devices, machine)
	}
	return devices, nil
}

func (k *kidsDB) addKid(name string) error {
	statement, err := k.kdb.Prepare("INSERT INTO kids (firstname) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	return err
}

func (k *kidsDB) allKids() ([]kid, error) {
	rows, err := k.kdb.Query("SELECT id, firstname FROM kids")
	if err != nil {
		return nil, err
	}
	var id int
	var firstname string
	var kids []kid
	for rows.Next() {
		err := rows.Scan(&id, &firstname)
		if err != nil {
			return kids, err
		}
		kid := kid{id, firstname}
		kids = append(kids, kid)
	}
	return kids, nil
}

func (k *kidsDB) addDevice(kidID string, mac string) error {
	statement, err := k.kdb.Prepare("INSERT INTO devices (mac, kidId) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(mac, kidID)
	return err
}

func (k *kidsDB) allDevices() ([]device, error) {
	rows, err := k.kdb.Query("SELECT id, name, mac FROM devices")
	if err != nil {
		return nil, err
	}
	var id int
	var name string
	var mac string
	var devices []device
	for rows.Next() {
		err := rows.Scan(&id, &name, &mac)
		if err != nil {
			return devices, err
		}
		machine := device{id, name, mac}
		devices = append(devices, machine)
	}
	return devices, nil
}
