package crud

import (
	"fmt"
	"testing"
)

func TestDosen(t *testing.T) {
	data := []*Dosen{
		{ID: "1", NamaDosen: "Aji"},
		{ID: "2", NamaDosen: "Ajo"},
		{ID: "3", NamaDosen: "Aje"},
		{ID: "4", NamaDosen: "Eji"},
		{ID: "5", NamaDosen: "Ejo"},
		{ID: "6", NamaDosen: "Eje"},
		{ID: "7", NamaDosen: "Dji"},
		{ID: "8", NamaDosen: "Djo"},
		{ID: "9", NamaDosen: "Dje"},
		{ID: "10", NamaDosen: "Oji"},
		{ID: "11", NamaDosen: "Ojo"},
		{ID: "12", NamaDosen: "Oje"},
		{ID: "13", NamaDosen: "Iji"},
		{ID: "14", NamaDosen: "Ijo"},
		{ID: "15", NamaDosen: "Ije"},
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
		if err := data[0].Update(db, "Aja"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		temp := &Dosen{ID: data[0].ID}
		if err := temp.Get(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets", func(t *testing.T) {
		limit, offset := GetLimitOffset(1, 10)
		sort := "asc"

		dosens, err := GetsDosen(db, limit, offset, sort)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(dosens)
	})

	t.Run("Test Delete", func(t *testing.T) {
		temp := &Dosen{ID: data[1].ID}
		if err := temp.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
}
