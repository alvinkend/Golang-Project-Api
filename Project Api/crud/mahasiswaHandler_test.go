package crud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMahasiswaHandler(t *testing.T) {
	db, err := Setup(dbName)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
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
		{NPM: "175", Nama: "Joko", Kelas: "4IA07"},
		{NPM: "246", Nama: "Abdul", Kelas: "4IA08"},
		{NPM: "337", Nama: "Herman", Kelas: "4IA09"},
		{NPM: "853", Nama: "Adrian", Kelas: "4IA010"},
	}

	webHandler := http.HandlerFunc(SS)
	RegisDB(db)

	t.Run("Testing Post Handler", func(t *testing.T) {
		for _, item := range data {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/ss/mahasiswa", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}

			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			got := Mahasiswa{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}

			fmt.Println(got)
		}

	})

	t.Run("Testing Put handler Mahasiswa", func(t *testing.T) {
		res := httptest.NewRecorder()
		change := map[string]interface{}{
			"nama":  "Ramadhani",
			"kelas": "4IA06",
		}

		jsonUpdate, err := json.MarshalIndent(change, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		urls := fmt.Sprintf("/api/ss/mahasiswa/%s", data[0].NPM)
		req, err := http.NewRequest(http.MethodPut, urls, bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := Mahasiswa{}
		fmt.Println(Mahasiswa{})
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		fmt.Println(got)
	})

	t.Run("Testing for Delete Hundler", func(t *testing.T) {
		res := httptest.NewRecorder()
		urls := fmt.Sprintf("/api/ss/mahasiswa/%s", data[0].NPM)
		req, err := http.NewRequest(http.MethodDelete, urls, nil)
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "true" {
			t.Fatal("Data tidak terhapus")
		}
	})
}
