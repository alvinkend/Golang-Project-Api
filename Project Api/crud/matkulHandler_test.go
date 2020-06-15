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

func TestMatkulHandler(t *testing.T) {
	db, err := Setup(dbName)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
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

	webHandler := http.HandlerFunc(SS)
	RegisDB(db)

	t.Run("Testing Post Handler", func(t *testing.T) {
		for _, item := range data {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/ss/matkul", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}

			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			got := Matkul{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}

			fmt.Println(got)
		}

	})

	t.Run("Testing Put handler Matkul", func(t *testing.T) {
		res := httptest.NewRecorder()
		change := map[string]interface{}{
			"namaMatkul": "PKN",
			"sks":        "3",
		}

		jsonUpdate, err := json.MarshalIndent(change, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		urls := fmt.Sprintf("/api/ss/matkul/%s", data[0].ID)
		req, err := http.NewRequest(http.MethodPut, urls, bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := Matkul{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		fmt.Println(got)
	})

	t.Run("Testing for Delete Handler", func(t *testing.T) {
		res := httptest.NewRecorder()
		urls := fmt.Sprintf("/api/ss/matkul/%s", data[0].ID)
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
