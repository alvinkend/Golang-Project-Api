package crud

import (
	"fmt"
	"testing"
)

func TestMatkul(t *testing.T) {
	data := []*Matkul{
		{ID: "1", NamaMatkul: "Grafkom", SKS: "2"},
		{ID: "2", NamaMatkul: "PKN", SKS: "3"},
		{ID: "3", NamaMatkul: "Fisika", SKS: "3"},
		{ID: "4", NamaMatkul: "Matematika", SKS: "2"},
		{ID: "5", NamaMatkul: "IPA", SKS: "3"},
		{ID: "6", NamaMatkul: "IPS", SKS: "3"},
		{ID: "7", NamaMatkul: "Agama", SKS: "2"},
		{ID: "8", NamaMatkul: "Alprog", SKS: "3"},
		{ID: "9", NamaMatkul: "Bahasa Indonesia", SKS: "3"},
		{ID: "10", NamaMatkul: "Bahasa Inggris", SKS: "2"},
		{ID: "11", NamaMatkul: "Bahasa Jerman", SKS: "3"},
		{ID: "12", NamaMatkul: "Rekom", SKS: "3"},
		{ID: "13", NamaMatkul: "RPL", SKS: "2"},
		{ID: "14", NamaMatkul: "PKN", SKS: "3"},
		{ID: "15", NamaMatkul: "Fisika", SKS: "3"},
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
		if err := data[0].Update(db, "Kimia", "3"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		temp := &Matkul{ID: data[1].ID}
		if err := temp.Get(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets", func(t *testing.T) {
		limit, offset := GetLimitOffset(1, 10)
		sort := "asc"

		matkuls, err := GetsMatkul(db, limit, offset, sort)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(matkuls)
	})

	t.Run("Test Delete", func(t *testing.T) {
		temp := &Matkul{ID: data[1].ID}
		if err := temp.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
}
