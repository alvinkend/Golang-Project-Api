package crud

import (
	"fmt"
	"testing"
)

func TestPegawai(t *testing.T) {
	data := []*Pegawai{
		{ID: "1", NamaPegawai: "Nogroho"},
		{ID: "2", NamaPegawai: "Roby"},
		{ID: "3", NamaPegawai: "Kendra"},
		{ID: "4", NamaPegawai: "Aji"},
		{ID: "5", NamaPegawai: "Pangestu"},
		{ID: "6", NamaPegawai: "Dodi"},
		{ID: "7", NamaPegawai: "Ade"},
		{ID: "8", NamaPegawai: "Farah"},
		{ID: "9", NamaPegawai: "Joko"},
		{ID: "10", NamaPegawai: "Aji"},
		{ID: "11", NamaPegawai: "Herman"},
		{ID: "12", NamaPegawai: "Joko"},
		{ID: "13", NamaPegawai: "Alvin"},
		{ID: "14", NamaPegawai: "Adrian"},
		{ID: "15", NamaPegawai: "Nugraha"},
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
		if err := data[0].Update(db, "Nugrohe"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		temp := &Pegawai{ID: data[0].ID}
		if err := temp.Get(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets", func(t *testing.T) {
		limit, offset := GetLimitOffset(1, 10)
		sort := "asc"

		pegawais, err := GetsPegawai(db, limit, offset, sort)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(pegawais)
	})

	t.Run("Test Delete", func(t *testing.T) {
		temp := &Pegawai{ID: data[1].ID}
		if err := temp.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
}
