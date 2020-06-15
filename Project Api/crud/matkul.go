package crud

import (
	"database/sql"
	"fmt"
)

//Matkul struct
type Matkul struct {
	ID         string `json:"id"`
	NamaMatkul string `json:"namaMatkul"`
	SKS        string `json:"sks"`
}

//TBMatkul membuat table Matkul
var TBMatkul = `
Create Table matkul
	(id serial PRIMARY KEY,
	namaMatkul VARCHAR(60),
	sks VARCHAR(2));`

//Insert func untuk insert data
func (mk *Matkul) Insert(db *sql.DB) error {
	if mk.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "INSERT INTO matkul Values($1, $2, $3)"
	_, err := db.Exec(query, &mk.ID, &mk.NamaMatkul, &mk.SKS)
	return err
}

//Delete func untuk mendelete data
func (mk *Matkul) Delete(db *sql.DB) error {
	if mk.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "DELETE FROM matkul where id = $1"
	_, err := db.Exec(query, &mk.ID)
	return err
}

//Get func untuk menampilkan satu data berdasarkan id
func (mk *Matkul) Get(db *sql.DB) error {
	if mk.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "Select * FROM matkul where id = $1"
	return db.QueryRow(query, &mk.ID).Scan(&mk.ID, &mk.NamaMatkul, &mk.SKS)
}

//Update func untuk update data
func (mk *Matkul) Update(db *sql.DB, namaMatkul string, sks string) error {
	if mk.ID == "" {
		return fmt.Errorf("id tidak bolehh kosong")
	}
	query := "UPDATE matkul Set namaMatkul = $1, sks = $2 WHERE id = $3"
	_, err := db.Exec(query, namaMatkul, sks, &mk.ID)
	return err
}

//GetsMatkul Menampilkan semua data pada tabel
func GetsMatkul(db *sql.DB, limit, offset int, sort string) ([]*Matkul, error) {
	query := "SELECT * FROM matkul"

	if sort != "" {
		query = fmt.Sprintf("%s ORDER BY namaMatkul %s", query, sort)
	}

	if limit > 0 || offset > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}

	result := []*Matkul{}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		matkul := &Matkul{}
		if err := rows.Scan(&matkul.ID, &matkul.NamaMatkul, &matkul.SKS); err != nil {
			return nil, err
		}
		result = append(result, matkul)
	}
	return result, err
}
