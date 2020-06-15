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

func TestPegawaiHandler(t *testing.T) {
	db, err := Setup(dbName)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
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

	webHandler := http.HandlerFunc(SS)
	RegisDB(db)

	t.Run("Testing Post Handler", func(t *testing.T) {
		for _, item := range data {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/ss/pegawai", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}

			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			got := Pegawai{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}

			fmt.Println(got)
		}

	})

	t.Run("Testing Put handler Pegawai", func(t *testing.T) {
		res := httptest.NewRecorder()
		change := map[string]interface{}{
			"namaPegawai": "Nugrohe",
		}

		jsonUpdate, err := json.MarshalIndent(change, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		urls := fmt.Sprintf("/api/ss/pegawai/%s", data[0].ID)
		req, err := http.NewRequest(http.MethodPut, urls, bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := Pegawai{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		fmt.Println(got)
	})

	t.Run("Testing for Delete Handler", func(t *testing.T) {
		res := httptest.NewRecorder()
		urls := fmt.Sprintf("/api/ss/pegawai/%s", data[0].ID)
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
