package crud

import (
	"database/sql"
	"fmt"

	//import lib/pq
	_ "github.com/lib/pq"
)

//Connect fungsi ini digunakan untuk melakukan koneksi dengan database
func Connect(name, password, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", name, password, dbName)

	db, err := sql.Open("postgres", connStr)
	err = db.Ping()
	return db, err
}

//CreateDB membuat database
func CreateDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("Create Database %s", name)
	_, err := db.Exec(query)
	return err
}

//DropDB mengahapus database
func DropDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("Drop Database if Exists %s", name)
	_, err := db.Exec(query)
	return err
}

//CreateTable membuat table database
func CreateTable(db *sql.DB, query string) error {

	_, err := db.Exec(query)
	return err
}

//Setup database
func Setup(dbName string) (*sql.DB, error) {
	//conntect ke database "username", "password", "dbname"

	db, err := Connect("postgres", "postgres", "postgres")
	if err != nil {
		return nil, err
	}

	if err := DropDB(db, dbName); err != nil {
		return nil, err
	}

	//create database
	if err := CreateDB(db, "databaseboongan"); err != nil {
		return nil, err
	}
	db, err = Connect("postgres", "postgres", dbName)
	if err != nil {
		return nil, err
	}
	if err := CreateTable(db, TBMahasiswa); err != nil {
		return nil, err
	}
	if err := CreateTable(db, TBNilai); err != nil {
		return nil, err
	}
	if err := CreateTable(db, TBDosen); err != nil {
		return nil, err
	}
	if err := CreateTable(db, TBMatkul); err != nil {
		return nil, err
	}
	if err := CreateTable(db, TBPegawai); err != nil {
		return nil, err
	}
	return db, nil
}
