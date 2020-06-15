package crud

import (
	"fmt"
	"testing"
)

func TestNilai(t *testing.T) {
	data := []*Nilai{
		{NPM: "121", NamaMHS: "Alvin", NamaMatkul: "IPA", Nilai: "90"},
		{NPM: "222", NamaMHS: "Roby", NamaMatkul: "IPS", Nilai: "80"},
		{NPM: "331", NamaMHS: "Kendra", NamaMatkul: "MTK", Nilai: "100"},
		{NPM: "151", NamaMHS: "Nugroho", NamaMatkul: "IPA", Nilai: "90"},
		{NPM: "252", NamaMHS: "Aji", NamaMatkul: "IPS", Nilai: "80"},
		{NPM: "313", NamaMHS: "Pangestu", NamaMatkul: "MTK", Nilai: "100"},
		{NPM: "171", NamaMHS: "Dodi", NamaMatkul: "IPA", Nilai: "90"},
		{NPM: "822", NamaMHS: "Budi", NamaMatkul: "IPS", Nilai: "80"},
		{NPM: "933", NamaMHS: "Ade", NamaMatkul: "MTK", Nilai: "100"},
		{NPM: "411", NamaMHS: "Kiki", NamaMatkul: "IPA", Nilai: "90"},
		{NPM: "522", NamaMHS: "Farah", NamaMatkul: "IPS", Nilai: "80"},
		{NPM: "233", NamaMHS: "Joko", NamaMatkul: "MTK", Nilai: "100"},
		{NPM: "671", NamaMHS: "Jaka", NamaMatkul: "IPA", Nilai: "90"},
		{NPM: "151", NamaMHS: "Herman", NamaMatkul: "IPS", Nilai: "80"},
		{NPM: "351", NamaMHS: "Adrian", NamaMatkul: "MTK", Nilai: "100"},
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
		temp := &Nilai{ID: "1"}
		if err := temp.Update(db, "44", "Ripin", "Grafkom", "80"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		temp := &Nilai{ID: "1"}
		if err := temp.Get(db); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test Gets", func(t *testing.T) {
		limit, offset := GetLimitOffset(1, 10)
		sort := "asc"

		nilais, err := GetsNilai(db, limit, offset, sort)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(nilais)
	})

	t.Run("Test Delete", func(t *testing.T) {
		temp := &Nilai{NPM: data[0].NPM}
		if err := temp.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
}
