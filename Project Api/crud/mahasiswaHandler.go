package crud

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

var (
	dbName = "databaseboongan"
)

//HandlerMahasiswaPost -> mengirim data ke database
func HandlerMahasiswaPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data Mahasiswa
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate
	// npm can only be a number
	// validate pas ngepost
	matchNPM, _ := regexp.MatchString("^[0-9]*$", data.NPM)
	if !matchNPM {
		http.Error(w, "npm wajib angka!", http.StatusBadRequest)
		return
	}

	matchNama, _ := regexp.MatchString("^[a-zA-Z- ]*$", data.Nama)
	if !matchNama {
		http.Error(w, "nama harus huruf!", http.StatusBadRequest)
		return
	}

	matchKelas, _ := regexp.MatchString("^[a-zA-Z0-9- ]*$", data.Kelas)
	if !matchKelas {
		http.Error(w, "kelas harus huruf dan angka!", http.StatusBadRequest)
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

//HandlerMahasiswaGet mengambil data satu mahasiswa
func HandlerMahasiswaGet(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)

	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	sort := r.URL.Query().Get("sort")

	// log.Println("page and size:", page, size)

	if last == "mahasiswa" {
		limit, offset := GetLimitOffset(int(page), int(size))

		data, err := GetsMahasiswa(db, limit, offset, sort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
	} else {
		temp := Mahasiswa{NPM: last}
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

//HandlerMahasiswaDelete untuk mengdelete data pada mahasiswa
func HandlerMahasiswaDelete(w http.ResponseWriter, r *http.Request) {
	last := LastIndex(r)
	data := Mahasiswa{NPM: last}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("true"))
}

//HandlerMahasiswaUpdate untuk mengupdate data pada mahasiswa
func HandlerMahasiswaUpdate(w http.ResponseWriter, r *http.Request) {
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
	data := Mahasiswa{NPM: last}

	// validate
	matchNama, _ := regexp.MatchString("^[a-zA-Z- ]*$", jsonMap["nama"].(string))
	if !matchNama {
		http.Error(w, "nama harus huruf!", http.StatusBadRequest)
		return
	}

	matchKelas, _ := regexp.MatchString("^[a-zA-Z0-9- ]*$", jsonMap["kelas"].(string))
	if !matchKelas {
		http.Error(w, "kelas hanya boleh huruf dan angka", http.StatusBadRequest)
		return
	}
	if err := data.Update(db, jsonMap["nama"].(string), jsonMap["kelas"].(string)); err != nil {
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
