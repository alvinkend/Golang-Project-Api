package crud

import (
	"database/sql"
	"fmt"
)

//Dosen struct
type Dosen struct {
	ID        string `json:"id"`
	NamaDosen string `json:"namaDosen"`
}

//TBDosen membuat table Dosen
var TBDosen = `
Create Table dosen
	(id VARCHAR PRIMARY KEY,
	namaDosen VARCHAR(60));`

//Insert func untuk insert data
func (d *Dosen) Insert(db *sql.DB) error {
	if d.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "INSERT INTO dosen Values($1, $2)"
	_, err := db.Exec(query, &d.ID, &d.NamaDosen)
	return err
}

//Delete func untuk mendelete data
func (d *Dosen) Delete(db *sql.DB) error {
	if d.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "DELETE FROM dosen where id = $1"
	_, err := db.Exec(query, &d.ID)
	return err
}

//Get func untuk menampilkan satu data berdasarkan id
func (d *Dosen) Get(db *sql.DB) error {
	if d.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "Select * FROM dosen where id = $1"
	return db.QueryRow(query, &d.ID).Scan(&d.ID, &d.NamaDosen)
}

//Update func untuk update data
func (d *Dosen) Update(db *sql.DB, namaDosen string) error {
	if d.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "UPDATE dosen Set namaDosen = $1 WHERE id = $2"
	_, err := db.Exec(query, namaDosen, &d.ID)
	return err
}

//GetsDosen Menampilkan Seluruh data pdaa tabel
func GetsDosen(db *sql.DB, limit, offset int, sort string) ([]*Dosen, error) {
	query := "SELECT * FROM dosen"

	if sort != "" {
		query = fmt.Sprintf("%s ORDER BY namaDosen %s", query, sort)
	}

	if limit > 0 || offset > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}

	result := []*Dosen{}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		dosen := &Dosen{}
		if err := rows.Scan(&dosen.ID, &dosen.NamaDosen); err != nil {
			return nil, err
		}
		result = append(result, dosen)
	}
	return result, err
}
