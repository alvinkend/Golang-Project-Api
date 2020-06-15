package crud

import (
	"fmt"

	"testing"
)

func TestMahasiswa(t *testing.T) {
	data := []*Mahasiswa{
		{NPM: "125", Nama: "Roby", Kelas: "4IA06"},
		{NPM: "231", Nama: "Alvin", Kelas: "4IA02"},
		{NPM: "353", Nama: "Kendra", Kelas: "4IA04"},
		{NPM: "463", Nama: "Nugroho", Kelas: "4IA03"},
		{NPM: "513", Nama: "Aji", Kelas: "4IA05"},
		{NPM: "636", Nama: "Pangestu", Kelas: "4IA08"},
		{NPM: "789", Nama: "Dodi", Kelas: "4IA04"},
		{NPM: "821", Nama: "Budi", Kelas: "4IA06"},
		{NPM: "991", Nama: "Ade", Kelas: "4IA03"},
		{NPM: "112", Nama: "Kiki", Kelas: "4IA04"},
		{NPM: "242", Nama: "Farah", Kelas: "4IA01"},
		{NPM: "366", Nama: "Evan", Kelas: "4IA08"},
		{NPM: "175", Nama: "Joko", Kelas: "4IA07"},
		{NPM: "246", Nama: "Abdul", Kelas: "4IA08"},
		{NPM: "337", Nama: "Herman", Kelas: "4IA09"},
		{NPM: "853", Nama: "Adrian", Kelas: "4IA010"},
	}

	db, err := Setup("databaseboongan")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Test Insert", func(t *testing.T) {
		for _, item := range data {
			if err := item.Insert(db); err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Test Update", func(t *testing.T) {
		if err := data[0].Update(db, "Ripin", "4IA06"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		temp := &Mahasiswa{NPM: data[0].NPM}
		if err := temp.Get(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets", func(t *testing.T) {
		limit, offset := GetLimitOffset(1, 10)
		sort := "asc"

		mahasiswas, err := GetsMahasiswa(db, limit, offset, sort)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(mahasiswas)
	})

	t.Run("Test Delete", func(t *testing.T) {
		temp := &Mahasiswa{NPM: data[0].NPM}
		if err := temp.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
}
