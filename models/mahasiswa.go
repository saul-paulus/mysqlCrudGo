package models

import "time"

type Mahasiswa struct {
	ID        int       `json:"id"`
	NIM       int       `json:"nim"`
	Nama      string    `json:"nama"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
