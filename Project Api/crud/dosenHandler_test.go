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

func TestDosenHandler(t *testing.T) {
	db, err := Setup(dbName)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
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

	webHandler := http.HandlerFunc(SS)
	RegisDB(db)

	t.Run("Testing Post Handler", func(t *testing.T) {
		for _, item := range data {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", " ")
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/ss/dosen", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}

			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			got := Dosen{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}

			fmt.Println(got)
		}

	})

	t.Run("Testing Put handler Dosen", func(t *testing.T) {
		res := httptest.NewRecorder()
		change := map[string]interface{}{
			"namaDosen": "Aja",
		}

		jsonUpdate, err := json.MarshalIndent(change, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		urls := fmt.Sprintf("/api/ss/dosen/%s", data[0].ID)
		req, err := http.NewRequest(http.MethodPut, urls, bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := Dosen{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		fmt.Println(got)
	})

	t.Run("Testing for Delete Handler", func(t *testing.T) {
		res := httptest.NewRecorder()
		urls := fmt.Sprintf("/api/ss/dosen/%s", data[0].ID)
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
