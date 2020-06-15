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

//HandlerMatkulPost -> mengirim data ke database
func HandlerMatkulPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Matkul
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate
	matchMatkul, _ := regexp.MatchString("^[a-zA-Z0-9- ]*$", data.NamaMatkul)
	if !matchMatkul {
		http.Error(w, "nama matkul harus huruf dan angka!", http.StatusBadRequest)
		return
	}

	matchSKS, _ := regexp.MatchString("^[0-9]*$", data.SKS)
	if !matchSKS {
		http.Error(w, "sks harus angka!", http.StatusBadRequest)
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

//HandlerMatkulGet mengambil data
func HandlerMatkulGet(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	sort := r.URL.Query().Get("sort")

	if last == "matkul" {

		limit, offset := GetLimitOffset(int(page), int(size))
		data, err := GetsMatkul(db, limit, offset, sort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		temp := Matkul{ID: last}
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

//HandlerMatkulDelete untuk mengdelete data pada mahasiswa
func HandlerMatkulDelete(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)
	data := Matkul{ID: last}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("true"))
}

//HandlerMatkulUpdate untuk mengupdate data pada mahasiswa
func HandlerMatkulUpdate(w http.ResponseWriter, r *http.Request) {
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
	data := Matkul{ID: last}

	// validate
	matchMatkul, _ := regexp.MatchString("^[a-zA-Z0-9- ]*$", jsonMap["namaMatkul"].(string))
	if !matchMatkul {
		http.Error(w, "nama matkul harus huruf dan angka!", http.StatusBadRequest)
		return
	}

	matchSKS, _ := regexp.MatchString("^[0-9]*$", jsonMap["sks"].(string))
	if !matchSKS {
		http.Error(w, "sks harus angka!", http.StatusBadRequest)
		return
	}

	if err := data.Update(db, jsonMap["namaMatkul"].(string), jsonMap["sks"].(string)); err != nil {
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
