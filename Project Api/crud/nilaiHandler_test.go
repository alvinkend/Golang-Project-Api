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

func TestNilaiHandler(t *testing.T) {
	db, err := Setup(dbName)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()
	data := []*Nilai{
		{ID: "1", NPM: "121", NamaMHS: "Alvin", NamaMatkul: "IPA", Nilai: "90"},
		{ID: "2", NPM: "222", NamaMHS: "Roby", NamaMatkul: "IPS", Nilai: "80"},
		{ID: "3", NPM: "331", NamaMHS: "Kendra", NamaMatkul: "MTK", Nilai: "100"},
		{ID: "4", NPM: "151", NamaMHS: "Nugroho", NamaMatkul: "IPA", Nilai: "90"},
		{ID: "5", NPM: "252", NamaMHS: "Aji", NamaMatkul: "IPS", Nilai: "80"},
		{ID: "6", NPM: "313", NamaMHS: "Pangestu", NamaMatkul: "MTK", Nilai: "100"},
		{ID: "7", NPM: "171", NamaMHS: "Dodi", NamaMatkul: "IPA", Nilai: "90"},
		{ID: "8", NPM: "822", NamaMHS: "Budi", NamaMatkul: "IPS", Nilai: "80"},
		{ID: "9", NPM: "933", NamaMHS: "Ade", NamaMatkul: "MTK", Nilai: "100"},
		{ID: "10", NPM: "411", NamaMHS: "Kiki", NamaMatkul: "IPA", Nilai: "90"},
		{ID: "11", NPM: "522", NamaMHS: "Farah", NamaMatkul: "IPS", Nilai: "80"},
		{ID: "12", NPM: "233", NamaMHS: "Joko", NamaMatkul: "MTK", Nilai: "100"},
		{ID: "13", NPM: "671", NamaMHS: "Jaka", NamaMatkul: "IPA", Nilai: "90"},
		{ID: "14", NPM: "151", NamaMHS: "Herman", NamaMatkul: "IPS", Nilai: "80"},
		{ID: "15", NPM: "351", NamaMHS: "Adrian", NamaMatkul: "MTK", Nilai: "100"},
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

			req, err := http.NewRequest(http.MethodPost, "/api/ss/nilai", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}

			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			got := Nilai{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}

			fmt.Println(got)
		}

	})

	t.Run("Testing Put handler Nilai", func(t *testing.T) {
		res := httptest.NewRecorder()
		change := map[string]interface{}{
			"npm":        "44",
			"namaMHS":    "Ripin",
			"namaMatkul": "Grafkom",
			"nilai":      "90",
		}

		jsonUpdate, err := json.MarshalIndent(change, "", " ")
		if err != nil {
			t.Fatal(err)
		}

		urls := fmt.Sprintf("/api/ss/nilai/%s", data[0].ID)
		req, err := http.NewRequest(http.MethodPut, urls, bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}

		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		got := Nilai{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		fmt.Println(got)
	})

	t.Run("Testing for Delete Handler", func(t *testing.T) {
		res := httptest.NewRecorder()
		urls := fmt.Sprintf("/api/ss/nilai/%s", data[0].ID)
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
