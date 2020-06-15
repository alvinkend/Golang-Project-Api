package crud

import (
	"database/sql"
	"fmt"
)

//Pegawai struct
type Pegawai struct {
	ID          string `json:"id"`
	NamaPegawai string `json:"namaPegawai"`
}

//TBPegawai membuat table pegawai
var TBPegawai = `
Create Table pegawai
	(id VARCHAR PRIMARY KEY,
	namaPegawai VARCHAR(60));`

//Insert func untuk insert data
func (pg *Pegawai) Insert(db *sql.DB) error {
	if pg.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "INSERT INTO pegawai Values($1, $2)"
	_, err := db.Exec(query, &pg.ID, &pg.NamaPegawai)
	return err
}

//Delete func untuk mendelete data
func (pg *Pegawai) Delete(db *sql.DB) error {
	if pg.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "DELETE FROM pegawai where id = $1"
	_, err := db.Exec(query, &pg.ID)
	return err
}

//Get func untuk menampilkan satu data berdasarkan id
func (pg *Pegawai) Get(db *sql.DB) error {
	if pg.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "Select * FROM pegawai where id = $1"
	return db.QueryRow(query, &pg.ID).Scan(&pg.ID, &pg.NamaPegawai)
}

//Update func untuk update data
func (pg *Pegawai) Update(db *sql.DB, namaPegawai string) error {
	if pg.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "UPDATE pegawai Set namaPegawai = $1 WHERE id = $2"
	_, err := db.Exec(query, namaPegawai, &pg.ID)
	return err
}

//GetsPegawai Menampilkan semua data pada tabel
func GetsPegawai(db *sql.DB, limit, offset int, sort string) ([]*Pegawai, error) {
	query := "SELECT * FROM pegawai"

	if sort != "" {
		query = fmt.Sprintf("%s ORDER BY namaPegawai %s", query, sort)
	}

	if limit > 0 || offset > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}
	result := []*Pegawai{}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pegawai := &Pegawai{}
		if err := rows.Scan(&pegawai.ID, &pegawai.NamaPegawai); err != nil {
			return nil, err
		}
		result = append(result, pegawai)
	}
	return result, err
}
