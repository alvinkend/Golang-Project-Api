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

//HandlerDosenPost -> mengirim data ke database
func HandlerDosenPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Dosen
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate
	// namaDosen can only be a character and space
	match, _ := regexp.MatchString("^[a-zA-Z- ]*$", data.NamaDosen)
	if !match {
		http.Error(w, "Nama Dosen Harus huruf", http.StatusBadRequest)
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

//HandlerDosenGet mengatur url get dosen
func HandlerDosenGet(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	sort := r.URL.Query().Get("sort")

	if last == "dosen" {
		limit, offset := GetLimitOffset(int(page), int(size))
		data, err := GetsDosen(db, limit, offset, sort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		temp := Dosen{ID: last}
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

//HandlerDosenDelete untuk mengdelete data pada mahasiswa
func HandlerDosenDelete(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)
	data := Dosen{ID: last}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("true"))
}

//HandlerDosenUpdate untuk mengupdate data pada mahasiswa
func HandlerDosenUpdate(w http.ResponseWriter, r *http.Request) {
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
	data := Dosen{ID: last}

	//validate
	matchNamaDosen, _ := regexp.MatchString("^[a-zA-Z- ]*$", jsonMap["namaDosen"].(string))
	if !matchNamaDosen {
		http.Error(w, "nama dosen harus huruf!", http.StatusBadRequest)
		return
	}
	if err := data.Update(db, jsonMap["namaDosen"].(string)); err != nil {
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
