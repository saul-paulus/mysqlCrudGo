package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pendalaman-res-api/config"
	"pendalaman-res-api/mahasiswa"
	"pendalaman-res-api/models"
	"pendalaman-res-api/utils"
	"strconv"
)

type Server struct {
	db      *sql.DB
	ViewDir string
}

func InitServer() *Server {
	dbCon, err := config.Mysql()
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		db:      dbCon,
		ViewDir: "/views",
	}

}

func (s *Server) Listen() {
	// Routing
	http.HandleFunc("/index", GetIndex)
	http.HandleFunc("/api/v1/getallmahasiswa", s.GetAllMahasiswa())
	http.HandleFunc("/api/v1/postmahasiswa", s.PostMahasiswa())
	http.HandleFunc("/api/v1/putmahasiswa", s.UpdateMahasiswa())
	http.HandleFunc("/api/v1/deletemahasiswa", s.DeleteMahasiswa())

	fmt.Println("Server berjalan di port 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	server := InitServer()
	server.Listen()

}

// GetIndex
func GetIndex(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJson(w, map[string]interface{}{
		"message": "selamat datang di server backend",
		"status":  200,
	})

}

// GetAllMahasiswa
func (s *Server) GetAllMahasiswa() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mahasiswas, err := mahasiswa.GetAll(s.db)
		if utils.IsError(w, err) {
			return
		}

		utils.ResponseJson(w, map[string]interface{}{
			"data": mahasiswas,
		})
		return
	}

}

// Create data Mahasiswa
func (s *Server) PostMahasiswa() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, s.ViewDir+"mahasiswa_create.html")
			return
		}

		mhs := models.Mahasiswa{}

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.IsError(w, err)
			return
		}

		if err := mahasiswa.CreateMhs(s.db, &mhs); err != nil {
			utils.IsError(w, err)
			return
		}

		utils.ResponseJson(w, map[string]interface{}{
			"data": mhs,
		})

		utils.ResponseJson(w, map[string]string{
			"message": "Data berhasil dibuat",
		})
	}
}

// Update data mahasiswa
func (s *Server) UpdateMahasiswa() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mhs := models.Mahasiswa{}

		if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
			utils.IsError(w, err)
			return
		}

		if err := mahasiswa.UpdateMhs(s.db, &mhs); err != nil {
			utils.IsError(w, err)
			return
		}

		utils.ResponseJson(w, map[string]string{
			"message": "Data berhasil di update",
		})
	}
}

//Delete data Mahasiswa
func (s *Server) DeleteMahasiswa() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		mhs := models.Mahasiswa{}

		id := r.URL.Query().Get("id")

		if id == "" {
			fmt.Println("id tidak boleh kosong")

		}

		mhs.ID, _ = strconv.Atoi(id)
		if err := mahasiswa.DeleteMhs(s.db, &mhs); err != nil {
			utils.IsError(w, err)
			return
		}

		utils.ResponseJson(w, map[string]interface{}{
			"message": "Data berhasil dihapus",
		})
	}
}
