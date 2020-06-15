package crud

import (
	"database/sql"
	"fmt"
)

//Nilai struct
type Nilai struct {
	ID         string `json:"id"`
	NPM        string `json:"npm"`
	NamaMHS    string `json:"namaMHS"`
	NamaMatkul string `json:"namaMatkul"`
	Nilai      string `json:"nilai"`
}

//TBNilai membuat table nilai
var TBNilai = `
Create Table nilai
	(id serial PRIMARY KEY,
	npm VARCHAR(10),
	namaMHS VARCHAR(60),
	namaMatkul VARCHAR(25),
	nilai VARCHAR(3));`

//Insert func untuk insert data
func (n *Nilai) Insert(db *sql.DB) error {
	query := "INSERT INTO nilai (npm, namaMHS, namaMatkul ,nilai) Values ($1, $2, $3, $4)"
	_, err := db.Exec(query, &n.NPM, &n.NamaMHS, &n.NamaMatkul, &n.Nilai)
	return err
}

//Delete func untuk mendelete data
func (n *Nilai) Delete(db *sql.DB) error {
	if n.NPM == "" {
		return fmt.Errorf("npm tidak bolehh kosong")
	}
	query := "DELETE FROM nilai where npm = $1"
	_, err := db.Exec(query, &n.NPM)
	return err
}

//Get func untuk menampilkan satu data berdasarkan id
func (n *Nilai) Get(db *sql.DB) error {
	if n.ID == "" {
		return fmt.Errorf("npm tidak bolehh kosong")
	}
	query := "Select * FROM nilai where id = $1"
	return db.QueryRow(query, &n.ID).Scan(&n.ID, &n.NPM, &n.NamaMHS, &n.NamaMatkul, &n.Nilai)
}

//Update func untuk update data
func (n *Nilai) Update(db *sql.DB, npm string, namaMHS string, namaMatkul string, nilai string) error {
	if n.ID == "" {
		return fmt.Errorf("ID tidak bolehh kosong")
	}
	query := "UPDATE nilai Set npm =$1, namaMHS = $2, namaMatkul = $3, nilai = $4 WHERE id = $5"
	_, err := db.Exec(query, npm, namaMHS, namaMatkul, nilai, &n.ID)
	return err
}

//GetsNilai Menampilkan semua data pada tabel
func GetsNilai(db *sql.DB, limit, offset int, sort string) ([]*Nilai, error) {
	query := "SELECT * FROM nilai"

	if sort != "" {
		query = fmt.Sprintf("%s ORDER BY namaMHS %s", query, sort)
	}

	if limit > 0 || offset > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}

	result := []*Nilai{}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		nilai := &Nilai{}
		if err := rows.Scan(&nilai.ID, &nilai.NPM, &nilai.NamaMHS, &nilai.NamaMatkul, &nilai.Nilai); err != nil {
			return nil, err
		}
		result = append(result, nilai)
	}
	return result, err
}
