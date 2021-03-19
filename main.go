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

type handler func(w http.ResponseWriter, r *http.Request)

// GetAllMahasiswa
func (s *Server) GetAllMahasiswa() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		mahasiswas, err := mahasiswa.GetAll(s.db)
		if utils.IsError(w, err) {
			return
		}

		utils.ResponseJson(w, map[string]interface{}{
			"data": mahasiswas,
		})
	}

}

// PostMahasiswa
func (s *Server) PostMahasiswa() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, s.ViewDir+"mahasiswa_create.html")
			return
		}
		mhs := models.Mahasiswa{}

		err := json.NewDecoder(r.Body).Decode(&mhs)

		if utils.IsError(w, err) {
			return
		}

		err = mahasiswa.CreateMhs(s.db, &mhs)
		if utils.IsError(w, err) {
			return
		}

		utils.ResponseJson(w, map[string]interface{}{
			"data": mhs,
		})
	}

}
