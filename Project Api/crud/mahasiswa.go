package crud

import (
	"database/sql"
	"fmt"
)

//Mahasiswa merupakan struct
type Mahasiswa struct {
	NPM   string `json:"npm"`
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
}

//TBMahasiswa membuat table mahasiswa
var TBMahasiswa = `
Create Table mahasiswa
	(npm VARCHAR(10) primary key,
	nama VARCHAR(60),
	kelas VARCHAR(10));`

//Insert func untuk insert data pada yabel
func (m *Mahasiswa) Insert(db *sql.DB) error {
	if m.NPM == "" {
		return fmt.Errorf("npm tidak bolehh kosong")
	}
	query := "INSERT INTO mahasiswa Values($1, $2, $3)"
	_, err := db.Exec(query, &m.NPM, &m.Nama, &m.Kelas)
	return err
}

//Delete func untuk mendelete data
func (m *Mahasiswa) Delete(db *sql.DB) error {
	if m.NPM == "" {
		return fmt.Errorf("npm tidak bolehh kosong")
	}
	query := "DELETE FROM mahasiswa where npm = $1"
	_, err := db.Exec(query, &m.NPM)
	return err
}

//Get func untuk menampilkan satu data berdasarkan npm
func (m *Mahasiswa) Get(db *sql.DB) error {
	if m.NPM == "" {
		return fmt.Errorf("npm tidak bolehh kosong")
	}
	query := "Select * FROM mahasiswa where npm = $1"
	return db.QueryRow(query, &m.NPM).Scan(&m.NPM, &m.Nama, &m.Kelas)
}

//Update func untuk update data
func (m *Mahasiswa) Update(db *sql.DB, nama string, kelas string) error {
	if m.NPM == "" {
		return fmt.Errorf("npm tidak bolehh kosong")
	}
	query := "UPDATE mahasiswa Set nama = $1, kelas = $2 WHERE npm = $3"
	_, err := db.Exec(query, nama, kelas, &m.NPM)
	return err
}

//GetsMahasiswa Menampilkan semua data pada tabel
func GetsMahasiswa(db *sql.DB, limit, offset int, sort string) ([]*Mahasiswa, error) {
	// log.Println("limit and offset:", limit, offset)

	query := "SELECT * FROM mahasiswa"

	if sort != "" {
		query = fmt.Sprintf("%s ORDER BY nama %s", query, sort)
	}

	if limit > 0 || offset > 0 {
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, limit, offset)
	}

	result := []*Mahasiswa{}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		mahasiswa := &Mahasiswa{}
		if err := rows.Scan(&mahasiswa.NPM, &mahasiswa.Nama, &mahasiswa.Kelas); err != nil {
			return nil, err
		}
		result = append(result, mahasiswa)
	}
	return result, err
}
