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

//HandlerPegawaiPost -> mengirim data ke database
func HandlerPegawaiPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Pegawai
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate
	match, _ := regexp.MatchString("^[a-zA-Z- ]*$", data.NamaPegawai)
	if !match {
		http.Error(w, "nama pegawai harus huruf!", http.StatusBadRequest)
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

//HandlerPegawaiGet mengambil data
func HandlerPegawaiGet(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	sort := r.URL.Query().Get("sort")

	if last == "pegawai" {
		limit, offset := GetLimitOffset(int(page), int(size))
		data, err := GetsPegawai(db, limit, offset, sort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		temp := Pegawai{ID: last}
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

//HandlerPegawaiDelete untuk mengdelete data pada mahasiswa
func HandlerPegawaiDelete(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)
	data := Pegawai{ID: last}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("true"))
}

//HandlerPegawaiUpdate untuk mengupdate data pada mahasiswa
func HandlerPegawaiUpdate(w http.ResponseWriter, r *http.Request) {
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
	data := Pegawai{ID: last}

	// validate
	match, _ := regexp.MatchString("^[a-zA-Z- ]*$", jsonMap["namaPegawai"].(string))
	if !match {
		http.Error(w, "nama pegawai harus huruf!", http.StatusBadRequest)
		return
	}

	if err := data.Update(db, jsonMap["namaPegawai"].(string)); err != nil {
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
