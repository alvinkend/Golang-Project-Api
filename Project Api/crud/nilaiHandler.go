package crud

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

// var (
// 	dbName = "databaseboongan"
// )

//HandlerNilaiPost -> mengirim data ke database
func HandlerNilaiPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Nilai
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate
	matchNPM, _ := regexp.MatchString("^[0-9]*$", data.NPM)
	if !matchNPM {
		http.Error(w, "npm harus angka!", http.StatusBadRequest)
		return
	}
	matchMHS, _ := regexp.MatchString("^[a-zA-Z- ]*$", data.NamaMHS)
	if !matchMHS {
		http.Error(w, "nama mahasiswa harus huruf!", http.StatusBadRequest)
		return
	}
	matchMatkul, _ := regexp.MatchString("^[a-zA-Z0-9- ]*$", data.NamaMatkul)
	if !matchMatkul {
		http.Error(w, "nama matkul harus huruf!", http.StatusBadRequest)
		return
	}
	matchNilai, _ := regexp.MatchString("^[0-9]*$", data.Nilai)
	if !matchNilai {
		http.Error(w, "nilai harus angka!", http.StatusBadRequest)
		return
	}

	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

//HandlerNilaiGet mengambil data
func HandlerNilaiGet(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	sort := r.URL.Query().Get("sort")

	if last == "nilai" {
		limit, offset := GetLimitOffset(int(page), int(size))
		data, err := GetsNilai(db, limit, offset, sort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		temp := Nilai{ID: last}
		if err := temp.Get(db); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(temp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)

	}
}

//HandlerNilaiDelete untuk mengdelete data pada mahasiswa
func HandlerNilaiDelete(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)
	data := Nilai{NPM: last}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("true"))
}

//HandlerNilaiUpdate untuk mengupdate data pada mahasiswa
func HandlerNilaiUpdate(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := Nilai{ID: last}

	// validate
	matchNPM, _ := regexp.MatchString("^[0-9]*$", jsonMap["npm"].(string))
	if !matchNPM {
		http.Error(w, "NPM harus angka!", http.StatusBadRequest)
		return
	}
	matchMHS, _ := regexp.MatchString("^[a-zA-Z- ]*$", jsonMap["namaMHS"].(string))
	if !matchMHS {
		http.Error(w, "nama harus huruf!", http.StatusBadRequest)
		return
	}
	matchMatkul, _ := regexp.MatchString("^[a-zA-Z0-9- ]*$", jsonMap["namaMatkul"].(string))
	if !matchMatkul {
		http.Error(w, "nama matkul harus huruf!", http.StatusBadRequest)
		return
	}
	matchNilai, _ := regexp.MatchString("^[0-9]*$", jsonMap["nilai"].(string))
	if !matchNilai {
		http.Error(w, "nilai harus angka!", http.StatusBadRequest)
		return
	}
	if err := data.Update(db, jsonMap["npm"].(string), jsonMap["namaMHS"].(string), jsonMap["namaMatkul"].(string), jsonMap["nilai"].(string)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := data.Get(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
