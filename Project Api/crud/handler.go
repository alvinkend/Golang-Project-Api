package crud

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

var db *sql.DB

//RegisDB belom tau
func RegisDB(sqlDB *sql.DB) {
	if db != nil {
		panic("db telah terdaftar")
	}
	db = sqlDB
}

//LastIndex untuk mendapatkan index terakir pada url
func LastIndex(r *http.Request) string {
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	lastIndex := dataURL[len(dataURL)-1]
	return lastIndex
}

//SS belum tau
func SS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8; application/json")
	dataURL := strings.Split(fmt.Sprintf("%s", r.URL.Path), "/")
	switch dataURL[3] {
	case "mahasiswa":
		switch r.Method {
		case http.MethodPost:
			HandlerMahasiswaPost(w, r)
		case http.MethodGet:
			HandlerMahasiswaGet(w, r)
		case http.MethodPut:
			HandlerMahasiswaUpdate(w, r)
		case http.MethodDelete:
			HandlerMahasiswaDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "nilai":
		switch r.Method {
		case http.MethodPost:
			HandlerNilaiPost(w, r)
		case http.MethodGet:
			HandlerNilaiGet(w, r)
		case http.MethodPut:
			HandlerNilaiUpdate(w, r)
		case http.MethodDelete:
			HandlerNilaiDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "dosen":
		switch r.Method {
		case http.MethodPost:
			HandlerDosenPost(w, r)
		case http.MethodGet:
			HandlerDosenGet(w, r)
		case http.MethodPut:
			HandlerDosenUpdate(w, r)
		case http.MethodDelete:
			HandlerDosenDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "matkul":
		switch r.Method {
		case http.MethodPost:
			HandlerMatkulPost(w, r)
		case http.MethodGet:
			HandlerMatkulGet(w, r)
		case http.MethodPut:
			HandlerMatkulUpdate(w, r)
		case http.MethodDelete:
			HandlerMatkulDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	case "pegawai":
		switch r.Method {
		case http.MethodPost:
			HandlerPegawaiPost(w, r)
		case http.MethodGet:
			HandlerPegawaiGet(w, r)
		case http.MethodPut:
			HandlerPegawaiUpdate(w, r)
		case http.MethodDelete:
			HandlerPegawaiDelete(w, r)
		default:
			w.Write([]byte("method tidak ditemukan"))
		}
	}
}
